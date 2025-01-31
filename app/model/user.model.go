package model

import (
	"github.com/bantawao4/gofiber-boilerplate/app/dao"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserModel struct {
	dao.User
}

func (u *UserModel) BeforeCreate(db *gorm.DB) error {

	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(password)
	if err != nil {
		return err
	}
	return nil
}
