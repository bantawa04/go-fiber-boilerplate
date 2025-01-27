package service

import (
	"fiber-boilerplate/app/model"
	"fiber-boilerplate/app/repository"
)

type UserService interface {
	GetUsers() ([]model.UserModel, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetUsers() ([]model.UserModel, error) {
	return s.userRepo.GetUsers()
}