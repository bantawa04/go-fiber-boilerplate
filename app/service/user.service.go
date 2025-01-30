package service

import (
	"github.com/bantawao4/gofiber-boilerplate/app/model"
	"github.com/bantawao4/gofiber-boilerplate/app/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUsers() ([]model.UserModel, error)
	CreateUser(user *model.UserModel) (*model.UserModel, error)
	ExistsByEmail(email string) bool
	ExistsByPhone(phone string) bool
}

// Add implementations for ExistsByEmail and ExistsByPhone
func (s *userService) ExistsByEmail(email string) bool {
	return s.userRepo.ExistsByEmail(email)
}

func (s *userService) ExistsByPhone(phone string) bool {
	return s.userRepo.ExistsByPhone(phone)
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Explicitly update the password
	user.Password = string(hashedPassword)

	// Save the user with the hashed password
	return s.userRepo.CreateUser(user)
}
