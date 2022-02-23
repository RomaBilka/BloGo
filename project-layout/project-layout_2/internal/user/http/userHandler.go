package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/RomaBiliak/BloGo/project-layout/project-layou_2/internal/user"
	"github.com/RomaBiliak/BloGo/project-layout/project-layou_2/pkg/response"
)

type userService interface {
	CreateUser(user user.User) (user.User, error)
	GetUser(id user.UserId) (user.User, error)
	DeleteUser(id user.UserId) error
	UpdateUser(id user.UserId, user user.User) (user.User, error)
	GetUsers() ([]user.User, error)
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
}

type createUserResponse struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *userHTTP) CreateUser(w http.ResponseWriter, r *http.Request) {
	userRequest := &createUserRequest{}

	err := json.NewDecoder(r.Body).Decode(userRequest)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	userModel := user.User{
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

func (h *userHTTP) UpdateUser(w http.ResponseWriter, r *http.Request) {
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

	userModel := user.User{
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

func (h *userHTTP) GetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	userInDb, err := h.userService.GetUser(userId)

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, createUserResponse{Id: uint64(userInDb.Id), Name: userInDb.Name, Email: userInDb.Email})
}

func (h *userHTTP) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetUsers()

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	var usersResponse []createUserResponse

	for _, u := range users {
		usersResponse = append(usersResponse, createUserResponse{Id: uint64(u.Id), Name: u.Name, Email: u.Email})
	}

	response.WriteJSON(w, http.StatusCreated, usersResponse)
}

func (h *userHTTP) DeleteUser(w http.ResponseWriter, r *http.Request) {
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

func getUserId(r *http.Request) (user.UserId, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		return 0, err
	}

	return user.UserId(id), nil
}
