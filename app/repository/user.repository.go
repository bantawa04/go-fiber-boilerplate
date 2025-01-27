package repository

import (
	"fiber-boilerplate/app/model"
	"fiber-boilerplate/config"
)

type UserRepository interface {
	GetUsers() ([]model.UserModel, error)
}

type userRepository struct {}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetUsers() ([]model.UserModel, error) {
	var users []model.UserModel
	err := config.DB.Db.Find(&users).Error
	return users, err
}