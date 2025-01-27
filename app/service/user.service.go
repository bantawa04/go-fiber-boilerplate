package service

import (
	"github.com/bantawao4/gofiber-boilerplate/app/model"
	"github.com/bantawao4/gofiber-boilerplate/app/repository"
)

type UserService interface {
	GetUsers() ([]model.UserModel, error)
	CreateUser(user *model.UserModel) (*model.UserModel, error)
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

func (s *userService) CreateUser(user *model.UserModel) (*model.UserModel, error) {
	return s.userRepo.CreateUser(user)
}
