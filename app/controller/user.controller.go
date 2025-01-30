package controller

import (
	"github.com/bantawao4/gofiber-boilerplate/app/request"
	"github.com/bantawao4/gofiber-boilerplate/app/service"
	"github.com/bantawao4/gofiber-boilerplate/app/validator"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetUsers(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
}

type userController struct {
	userService service.UserService
	validator   validator.Validator
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
		validator:   validator.NewValidator(),
	}
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

func (ctrl *userController) CreateUser(c *fiber.Ctx) error {
	reqData := new(request.CreateUserRequestData)

	if err := c.BodyParser(reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to bind user data",
		})
	}

	if errors := ctrl.validator.ValidateStruct(reqData); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors":  errors,
			"message": "Invalid input information",
		})
	}

	// Check if email exists
	if exists := ctrl.userService.ExistsByEmail(reqData.Email); exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Email already exists",
			"message": "User with this email already exists",
		})
	}

	// Check if phone exists
	if exists := ctrl.userService.ExistsByPhone(reqData.Phone); exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Phone already exists",
			"message": "User with this phone already exists",
		})
	}

	createdUser, err := ctrl.userService.CreateUser(&reqData.UserModel)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User Created Successfully",
		"data":    createdUser,
	})
}
