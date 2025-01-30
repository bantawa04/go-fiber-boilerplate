package repository

import (
	"github.com/bantawao4/gofiber-boilerplate/app/dao"
	"github.com/bantawao4/gofiber-boilerplate/app/model"
	"github.com/bantawao4/gofiber-boilerplate/config"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(page, perPage int) ([]dao.User, int64, error)
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

func (r *userRepository) GetUsers(page, perPage int) ([]dao.User, int64, error) {
	var users []dao.User
	var total int64

	// Get total count
	if err := r.db.Table("users").Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated data
	offset := (page - 1) * perPage
	err := r.db.Table("users").Offset(offset).Limit(perPage).Scan(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, err
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
