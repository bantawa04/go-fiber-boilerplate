package request

import "github.com/bantawao4/gofiber-boilerplate/app/model"

type CreateUserRequestData struct {
	model.UserModel
	// ConfirmPassword string `json:"confirm_password" validate:"required"`
}
