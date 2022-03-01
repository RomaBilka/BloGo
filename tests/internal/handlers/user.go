package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/RomaBiliak/BloGo/project-layout/tests/internal/models"
	"github.com/RomaBiliak/BloGo/project-layout/tests/pkg/response"
)

type userService interface {
	CreateUser(user models.User) (models.User, error)
}

type userHTTP struct {
	userService userService
}

func NewUserHttp(userService userService) *userHTTP {
	return &userHTTP{userService: userService}
}

type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (r *createUserRequest) validate() (bool, error) {
	if len(r.Name) < 3 {
		return false, fmt.Errorf("%s", "Bad request, short user name")
	}

	if len(r.Phone) < 9 {
		return false, fmt.Errorf("%s", "Bad request, short user phone")
	}

	_, err := mail.ParseAddress(r.Email)
	if err != nil {
		return false, fmt.Errorf("%s", "Bad request, wrong user email")
	}

	return true, nil
}

type createUserResponse struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (h *userHTTP) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.WriteERROR(w, http.StatusMethodNotAllowed, nil)
		return
	}

	userRequest := &createUserRequest{}

	err := json.NewDecoder(r.Body).Decode(userRequest)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	ok, err := userRequest.validate()
	if !ok {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	userModel := models.User{
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Phone: userRequest.Phone,
	}

	newUser, err := h.userService.CreateUser(userModel)

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, createUserResponse{Id: uint64(newUser.Id), Name: newUser.Name, Email: newUser.Email, Phone: newUser.Phone})
}
