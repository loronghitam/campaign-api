package usecase

import (
	"campaign-api/domain"
	"campaign-api/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (domain.User, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (domain.User, error) {
	user := domain.User{}
	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email
	Password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(Password)
	user.Role = "user"
	newUser, err := s.repository.SAVE(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
