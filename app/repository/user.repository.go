package repository

import (
	"github.com/bantawao4/gofiber-boilerplate/app/model"
	"github.com/bantawao4/gofiber-boilerplate/config"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(page, perPage int) ([]model.UserModel, int64, error)
	CreateUser(user *model.UserModel) (*model.UserModel, error)
	GetUserById(userId string) (*model.UserModel, error)
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

func (r *userRepository) GetUsers(page, perPage int) ([]model.UserModel, int64, error) {
	var users []model.UserModel
	var total int64

	// Get total count
	if err := r.db.Model(&model.UserModel{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated data
	offset := (page - 1) * perPage
	err := r.db.Model(&model.UserModel{}).Offset(offset).Limit(perPage).Scan(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, err
}

func (r *userRepository) CreateUser(user *model.UserModel) (*model.UserModel, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) ExistsByEmail(email string) bool {
	var count int64
	r.db.Model(&model.UserModel{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func (r *userRepository) ExistsByPhone(phone string) bool {
	var count int64
	r.db.Model(&model.UserModel{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}
func (r *userRepository) GetUserById(userId string) (*model.UserModel, error) {
	var user model.UserModel
	err := r.db.Model(&model.UserModel{}).Where("id = ?", userId).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
