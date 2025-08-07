package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) RegisterUser(username, email, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := s.repository.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(username, password string) (*User, error) {
	userFound, err := s.repository.FindByEmail(username)
	if err != nil {
		return nil, errors.New("invalid username")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return userFound, nil
}
