package services

import (
	"fmt"

	"github.com/RomaBilka/BloGo/project-layout/project-layou_1/internal/models"
)

type userRepository interface {
	CreateUser(user models.User) (models.UserId, error)
	GetUserById(id models.UserId) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	UpdateUser(id models.UserId, user models.User) error
	DeleteUser(id models.UserId) error
	GetUsers() ([]models.User, error)
}

type UserService struct {
	repository userRepository
}

func NewUserService(repository userRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	exists, err := s.repository.GetUserByEmail(user.Email)
	if err != nil {
		return user, err
	}

	if exists.Id > 0 {
		return models.User{}, fmt.Errorf("%s", "User with that email already exists")
	}

	id, err := s.repository.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}

	user, err = s.repository.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *UserService) GetUser(id models.UserId) (models.User, error) {
	user, err := s.repository.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}

	if user.Id == 0 {
		return models.User{}, fmt.Errorf("%s", "Bad request, user not found")
	}

	return user, nil
}

func (s *UserService) DeleteUser(id models.UserId) error {
	err := s.repository.DeleteUser(id)

	return err
}

func (s *UserService) UpdateUser(id models.UserId, user models.User) (models.User, error) {
	exists, err := s.repository.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}
	if exists.Id == 0 {
		return models.User{}, fmt.Errorf("%s", "Bad request, user not found")
	}

	exists, err = s.repository.GetUserByEmail(user.Email)
	if err != nil {
		return user, err
	}

	if exists.Id > 0 && exists.Id != id {
		return models.User{}, fmt.Errorf("%s", "User with that email already exists")
	}

	err = s.repository.UpdateUser(id, user)
	if err != nil {
		return models.User{}, err
	}

	user, err = s.repository.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}

func (s *UserService) GetUsers() ([]models.User, error) {
	users, err := s.repository.GetUsers()
	if err != nil {
		return []models.User{}, err
	}

	return users, nil
}
