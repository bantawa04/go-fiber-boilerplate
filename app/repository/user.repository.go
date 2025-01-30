package repository

import (
	"github.com/bantawao4/gofiber-boilerplate/app/model"
	"github.com/bantawao4/gofiber-boilerplate/config"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]model.UserModel, error)
	CreateUser(user *model.UserModel) (*model.UserModel, error)
	ExistsByEmail(email string) bool
	ExistsByPhone(phone string) bool
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: config.DB.Db,
	}
}

func (r *userRepository) GetUsers() ([]model.UserModel, error) {
	var users []model.UserModel
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) CreateUser(user *model.UserModel) (*model.UserModel, error) {
	err := r.db.Table("users").Create(&user.User).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) ExistsByEmail(email string) bool {
	var count int64
	r.db.Table("users").Where("email = ?", email).Count(&count)
	return count > 0
}

func (r *userRepository) ExistsByPhone(phone string) bool {
	var count int64
	r.db.Table("users").Where("phone = ?", phone).Count(&count)
	return count > 0
}
