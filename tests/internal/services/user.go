package services

import (
	"fmt"

	"github.com/RomaBilka/BloGo/tests/internal/models"
)

type userRepository interface {
	CheckUserExists(user models.User) (bool, error)
	GetUserById(id models.UserId) (models.User, error)
	CreateUser(user models.User) (models.UserId, error)
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
	exists, err := s.repository.CheckUserExists(user)
	if err != nil {
		return user, err
	}

	if exists {
		return models.User{}, fmt.Errorf("%s", "User with that phone or email already exists")
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
