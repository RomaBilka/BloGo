package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RomaBilka/BloGo/example-wire/internal/models"
	"github.com/RomaBilka/BloGo/example-wire/pkg/response"
	"github.com/gorilla/mux"
)

type UserService interface {
	CreateUser(user models.User) (models.User, error)
	GetUser(id models.UserId) (models.User, error)
	DeleteUser(id models.UserId) error
	UpdateUser(id models.UserId, user models.User) (models.User, error)
	GetUsers() ([]models.User, error)
}

type UserHTTP struct {
	userService UserService
}

func NewUserHttp(userService UserService) *UserHTTP {
	return &UserHTTP{userService: userService}
}

type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type createUserResponse struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *UserHTTP) CreateUser(w http.ResponseWriter, r *http.Request) {
	userRequest := &createUserRequest{}

	err := json.NewDecoder(r.Body).Decode(userRequest)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	userModel := models.User{
		Name:  userRequest.Name,
		Email: userRequest.Email,
	}

	newUser, err := h.userService.CreateUser(userModel)

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, createUserResponse{Id: uint64(newUser.Id), Name: newUser.Name, Email: newUser.Email})
}

func (h *UserHTTP) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	userRequest := &createUserRequest{}

	err = json.NewDecoder(r.Body).Decode(userRequest)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	userModel := models.User{
		Name:  userRequest.Name,
		Email: userRequest.Email,
	}

	user, err := h.userService.UpdateUser(userId, userModel)

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, createUserResponse{Id: uint64(user.Id), Name: user.Name, Email: user.Email})
}

func (h *UserHTTP) GetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.userService.GetUser(userId)

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, createUserResponse{Id: uint64(user.Id), Name: user.Name, Email: user.Email})
}

func (h *UserHTTP) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetUsers()

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	var usersResponse []createUserResponse

	for _, user := range users {
		usersResponse = append(usersResponse, createUserResponse{Id: uint64(user.Id), Name: user.Name, Email: user.Email})
	}

	response.WriteJSON(w, http.StatusCreated, usersResponse)
}

func (h *UserHTTP) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	err = h.userService.DeleteUser(userId)

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	response.WriteJSON(w, http.StatusNoContent, createUserResponse{})
}

func getUserId(r *http.Request) (models.UserId, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		return 0, err
	}

	return models.UserId(id), nil
}
