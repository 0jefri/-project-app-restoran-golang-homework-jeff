package service

import (
	"errors"

	"github.com/project-aplikasi-restorant/model"
	"github.com/project-aplikasi-restorant/repository"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func (s *UserService) Login(username, password string) (model.User, error) {
	user, err := s.UserRepo.GetUserByUsername(username)
	if err != nil {
		return model.User{}, errors.New("username tidak ditemukan")
	}
	if user.Password != password {
		return model.User{}, errors.New("password salah")
	}
	return user, nil
}
