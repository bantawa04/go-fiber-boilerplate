package controller

import (
	"strconv"

	"github.com/bantawao4/gofiber-boilerplate/app/dto"
	"github.com/bantawao4/gofiber-boilerplate/app/request"
	"github.com/bantawao4/gofiber-boilerplate/app/response"
	"github.com/bantawao4/gofiber-boilerplate/app/service"
	"github.com/bantawao4/gofiber-boilerplate/app/validator"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetUsers(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
}

type userController struct {
	userService service.UserService
	validator   validator.UserValidator
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
		validator:   validator.NewUserValidator(),
	}
}

func (ctrl *userController) GetUsers(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	perPage, err := strconv.Atoi(c.Query("perPage", "10"))
	if err != nil || perPage < 1 {
		perPage = 10
	}

	searchQuery := c.Query("search", "")

	users, meta, err := ctrl.userService.GetUsers(page, perPage, searchQuery)
	if err != nil {
		return err
	}

	return response.SuccessPaginationResponse(c, fiber.StatusOK, dto.ToUserListResponse(users), meta, "Users fetched successfully")
}

func (ctrl *userController) CreateUser(c *fiber.Ctx) error {
	reqData := new(request.CreateUserRequestData)

	if err := c.BodyParser(reqData); err != nil {
		return err
	}

	if errors := ctrl.validator.Validate.Struct(reqData); errors != nil {
		return response.ValidationErrorResponse(c,
			ctrl.validator.GenerateValidationResponse(errors))
	}

	userModel := reqData.ToModel()
	createdUser, err := ctrl.userService.CreateUser(userModel)
	if err != nil {
		return err
	}

	return response.SuccessDataResponse(c, fiber.StatusCreated,
		dto.ToUserResponse(createdUser), "User Created Successfully")
}

func (ctrl *userController) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := ctrl.userService.GetUserById(id)
	if err != nil {
		return err
	}

	return response.SuccessDataResponse(c, fiber.StatusOK,
		dto.ToUserResponse(user), "User fetched successfully")
}

func (ctrl *userController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	// Parse and validate request data
	reqData := new(request.UpdateUserRequestData)
	if err := c.BodyParser(reqData); err != nil {
		return err
	}

	if errors := ctrl.validator.Validate.Struct(reqData); errors != nil {
		return response.ValidationErrorResponse(c,
			ctrl.validator.GenerateValidationResponse(errors))
	}

	// Update user data from request
	updatedUser, err := ctrl.userService.UpdateUser(id, reqData.ToModel())
	if err != nil {
		return err
	}

	return response.SuccessDataResponse(c, fiber.StatusOK,
		dto.ToUserResponse(updatedUser), "User updated successfully")
}
