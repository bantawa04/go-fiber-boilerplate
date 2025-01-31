package dto

import (
	"github.com/bantawao4/gofiber-boilerplate/app/model"
	"time"
)

type UserResponse struct {
	ID        string     `json:"id"`
	FullName  string     `json:"full_name"`
	Phone     string     `json:"phone"`
	Gender    string     `json:"gender"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func ToUserResponse(user *model.UserModel) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Phone:     user.Phone,
		Gender:    user.Gender,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserListResponse(users []model.UserModel) []UserResponse {
	var response []UserResponse
	for _, user := range users {
		response = append(response, *ToUserResponse(&user))
	}
	return response
}