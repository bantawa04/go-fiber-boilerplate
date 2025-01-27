package controller

import (
	"fiber-boilerplate/app/service"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetUsers(c *fiber.Ctx) error
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{userService: userService}
}

func (ctrl *userController) GetUsers(c *fiber.Ctx) error {
	users, err := ctrl.userService.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": users,
	})
}