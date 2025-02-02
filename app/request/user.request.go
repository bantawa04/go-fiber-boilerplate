package request

import (
	"github.com/bantawao4/gofiber-boilerplate/app/dao"
	"github.com/bantawao4/gofiber-boilerplate/app/model"
)

type CreateUserRequestData struct {
	FullName string `json:"full_name" validate:"required"`
	Gender   string `json:"gender" validate:"required,gender"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Phone    string `json:"phone" validate:"required,phone"`
}

type UpdateUserRequestData struct {
	FullName string `json:"full_name"`
	Gender   string `json:"gender" validate:"gender"`
	Email    string `json:"email" validate:"email"`
	Phone    string `json:"phone" validate:"phone"`
}

// ToModel converts the request data to a UserModel
func (r *CreateUserRequestData) ToModel() *model.UserModel {
	return &model.UserModel{
		User: dao.User{
			FullName: r.FullName,
			Gender:   r.Gender,
			Email:    r.Email,
			Password: r.Password,
			Phone:    r.Phone,
		},
	}
}

// Add this method after CreateUserRequestData's ToModel
func (r *UpdateUserRequestData) ToModel() *model.UserModel {
    return &model.UserModel{
        User: dao.User{
            FullName: r.FullName,
            Gender:   r.Gender,
            Email:    r.Email,
            Phone:    r.Phone,
        },
    }
}
