package service

import (
	"math"

	"github.com/bantawao4/gofiber-boilerplate/app/model"
	"github.com/bantawao4/gofiber-boilerplate/app/repository"
	"github.com/bantawao4/gofiber-boilerplate/app/response"
)

type UserService interface {
	GetUsers(page, perPage int, searchQuery string) ([]model.UserModel, *response.PaginationMeta, error)
	CreateUser(user *model.UserModel) (*model.UserModel, error)
	ExistsByEmail(email string) bool
	ExistsByPhone(phone string) bool
	GetUserById(id string) (*model.UserModel, error)
}

type userService struct {
	userRepo repository.UserRepository
}

// Add implementations for ExistsByEmail and ExistsByPhone
func (s *userService) ExistsByEmail(email string) bool {
	return s.userRepo.ExistsByEmail(email)
}

func (s *userService) ExistsByPhone(phone string) bool {
	return s.userRepo.ExistsByPhone(phone)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetUsers(page, perPage int, searchQuery string) ([]model.UserModel, *response.PaginationMeta, error) {
	users, total, err := s.userRepo.GetUsers(page, perPage, searchQuery)
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
	return s.userRepo.CreateUser(user)
}

func (s *userService) GetUserById(id string) (*model.UserModel, error) {
	return s.userRepo.GetUserById(id)
}
