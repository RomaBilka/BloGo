package user

import (
	"fmt"
)

type userRepository interface {
	CreateUser(user User) (UserId, error)
	GetUserById(id UserId) (User, error)
	GetUserByEmail(email string) (User, error)
	UpdateUser(id UserId, user User) error
	DeleteUser(id UserId) error
	GetUsers() ([]User, error)
}

type UserService struct {
	repository userRepository
}

func NewUserService(repository userRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) CreateUser(user User) (User, error) {
	exists, err := s.repository.GetUserByEmail(user.Email)
	if err != nil {
		return user, err
	}

	if exists.Id > 0 {
		return User{}, fmt.Errorf("%s", "User with that email already exists")
	}

	id, err := s.repository.CreateUser(user)
	if err != nil {
		return User{}, err
	}

	user, err = s.repository.GetUserById(id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *UserService) GetUser(id UserId) (User, error) {
	user, err := s.repository.GetUserById(id)
	if err != nil {
		return User{}, err
	}

	if user.Id == 0 {
		return User{}, fmt.Errorf("%s", "Bad request, user not found")
	}

	return user, nil
}

func (s *UserService) DeleteUser(id UserId) error {
	err := s.repository.DeleteUser(id)

	return err
}

func (s *UserService) UpdateUser(id UserId, user User) (User, error) {
	exists, err := s.repository.GetUserById(id)
	if err != nil {
		return User{}, err
	}
	if exists.Id == 0 {
		return User{}, fmt.Errorf("%s", "Bad request, user not found")
	}

	exists, err = s.repository.GetUserByEmail(user.Email)
	if err != nil {
		return user, err
	}

	if exists.Id > 0 && exists.Id != id {
		return User{}, fmt.Errorf("%s", "User with that email already exists")
	}

	err = s.repository.UpdateUser(id, user)
	if err != nil {
		return User{}, err
	}

	user, err = s.repository.GetUserById(id)
	if err != nil {
		return User{}, err
	}

	return user, err
}

func (s *UserService) GetUsers() ([]User, error) {
	users, err := s.repository.GetUsers()
	if err != nil {
		return []User{}, err
	}

	return users, nil
}
