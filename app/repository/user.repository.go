package repository

import (
	"github.com/bantawao4/gofiber-boilerplate/app/model"
	"github.com/bantawao4/gofiber-boilerplate/config"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(page, perPage int, searchQuery string) ([]model.UserModel, int64, error)
	CreateUser(user *model.UserModel) (*model.UserModel, error)
	GetUserById(userId string) (*model.UserModel, error)
	GetUserByEmail(email string) bool
	GetUserByPhone(phone string) bool
	UpdateUser(user *model.UserModel) (*model.UserModel, error)
	DeleteUser(id string) error
	WithTrx(tx *gorm.DB) UserRepository
}

func (r *userRepository) WithTrx(tx *gorm.DB) UserRepository {
	if tx == nil {
		return r
	}
	// Create a new instance instead of modifying the existing one
	newRepo := &userRepository{
		db: tx,
	}
	return newRepo
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: config.DB.Db,
	}
}

func (r *userRepository) GetUsers(page, perPage int, searchQuery string) ([]model.UserModel, int64, error) {
	var users []model.UserModel
	var total int64

	query := r.db.Model(&model.UserModel{})

	if searchQuery != "" {
		query = query.Where("full_name ILIKE ? OR email ILIKE ?",
			"%"+searchQuery+"%",
			"%"+searchQuery+"%")
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated data
	offset := (page - 1) * perPage
	if err := query.Offset(offset).Limit(perPage).Scan(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepository) CreateUser(user *model.UserModel) (*model.UserModel, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
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

func (r *userRepository) UpdateUser(user *model.UserModel) (*model.UserModel, error) {
	err := r.db.Model(&model.UserModel{}).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) bool {
	var count int64
	r.db.Model(&model.UserModel{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func (r *userRepository) GetUserByPhone(phone string) bool {
	var count int64
	r.db.Model(&model.UserModel{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// Update implementation
func (r *userRepository) DeleteUser(id string) error {
	return r.db.Delete(&model.UserModel{}, "id = ?", id).Error
}
