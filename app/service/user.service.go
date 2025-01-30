package service

import (
	"math"

	"github.com/bantawao4/gofiber-boilerplate/app/dao"
	"github.com/bantawao4/gofiber-boilerplate/app/model"
	"github.com/bantawao4/gofiber-boilerplate/app/repository"
	"github.com/bantawao4/gofiber-boilerplate/app/response"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUsers(page, perPage int) ([]dao.User, *response.PaginationMeta, error)
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

func (s *userService) GetUsers(page, perPage int) ([]dao.User, *response.PaginationMeta, error) {
	users, total, err := s.userRepo.GetUsers(page, perPage)
	if err != nil {
		return nil, nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(perPage)))
	meta := &response.PaginationMeta{
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
		TotalItems: total,
	}

	return users, meta, nil
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
