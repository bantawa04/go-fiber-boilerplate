package service

import (
	"math"

	"github.com/bantawao4/gofiber-boilerplate/app/errors"
	"github.com/bantawao4/gofiber-boilerplate/app/model"
	"github.com/bantawao4/gofiber-boilerplate/app/repository"
	"github.com/bantawao4/gofiber-boilerplate/app/response"
)

// Update the interface definition
type UserService interface {
	GetUsers(page, perPage int, searchQuery string) ([]model.UserModel, *response.PaginationMeta, error)
	CreateUser(user *model.UserModel) (*model.UserModel, error)
	ExistsByEmail(email string) bool
	ExistsByPhone(phone string) bool
	GetUserById(id string) (*model.UserModel, error)
	UpdateUser(id string, user *model.UserModel) (*model.UserModel, error) // Changed signature
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

// Update the implementation to match interface
func (s *userService) UpdateUser(id string, updateData *model.UserModel) (*model.UserModel, error) {
	existingUser, err := s.userRepo.GetUserById(id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	if existingUser == nil {
		return nil, errors.NewNotFoundError("User not found")
	}

	if updateData.Email != "" && updateData.Email != existingUser.Email {
		if s.ExistsByEmail(updateData.Email) {
			return nil, errors.NewConflictError("Email already in use")
		}
	}

	// Update only the fields that are provided
	if updateData.FullName != "" {
		existingUser.FullName = updateData.FullName
	}
	if updateData.Phone != "" {
		existingUser.Phone = updateData.Phone
	}
	if updateData.Gender != "" {
		existingUser.Gender = updateData.Gender
	}
	if updateData.Email != "" {
		existingUser.Email = updateData.Email
	}

	return s.userRepo.UpdateUser(existingUser)
}
