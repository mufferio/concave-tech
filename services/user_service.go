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
	entities, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	dtosOut := make([]dtos.UserDTO, len(entities))
	for i, e := range entities {
		dtosOut[i] = dtos.UserDTO{
			Firstname:   e.Firstname,
			Lastname:    e.Lastname,
			DateOfBirth: e.DateOfBirth.Format("02-Jan-2006"),
			Income:      e.Income,
		}
	}
	return dtosOut, nil
}
