package services

import (
	"github.com/mufferio/concave-tech/dtos"
	"github.com/mufferio/concave-tech/repositories"
)

type UserService interface {
	GetUsers() ([]dtos.UserDTO, error)
}

type userService struct{ repo repositories.UserRepository }

func NewUserService(r repositories.UserRepository) UserService {
	return &userService{r}
}

func (s *userService) GetUsers() ([]dtos.UserDTO, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	out := make([]dtos.UserDTO, len(users))
	for i, u := range users {
		out[i] = dtos.UserDTO{
			Firstname:   u.Firstname,
			Lastname:    u.Lastname,
			DateOfBirth: u.DateOfBirth.Format("02-Jan-2006"),
			Income:      u.Income,
		}
	}
	return out, nil
}
