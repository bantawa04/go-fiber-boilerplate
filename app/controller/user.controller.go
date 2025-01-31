package controller

import (
	"fmt"
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
	// Get page from query, default to 1
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	// Get perPage from query, default to 10
	perPage, err := strconv.Atoi(c.Query("perPage", "10"))
	if err != nil || perPage < 1 {
		perPage = 10
	}
	searchQuery := c.Query("search","")

	users, meta, err := ctrl.userService.GetUsers(page, perPage, searchQuery)
	if err != nil {
		return response.ErrorResponse(c, fiber.StatusInternalServerError, err, "Failed to fetch users")
	}

	return response.SuccessPaginationResponse(c, fiber.StatusOK, dto.ToUserListResponse(users), meta, "Users fetched successfully")
}

func (ctrl *userController) CreateUser(c *fiber.Ctx) error {
	reqData := new(request.CreateUserRequestData)

	if err := c.BodyParser(reqData); err != nil {
		return response.ErrorResponse(c, fiber.StatusBadRequest, err, "Failed to bind user data")
	}

	if errors := ctrl.validator.Validate.Struct(reqData); errors != nil {
		return response.ValidationErrorResponse(c, ctrl.validator.GenerateValidationResponse(errors))
	}

	if exists := ctrl.userService.ExistsByEmail(reqData.Email); exists {
		return response.ErrorResponse(c, fiber.StatusBadRequest, fmt.Errorf("email already exists"), "User with this email already exists")
	}

	if exists := ctrl.userService.ExistsByPhone(reqData.Phone); exists {
		return response.ErrorResponse(c, fiber.StatusBadRequest, fmt.Errorf("phone already exists"), "User with this phone already exists")
	}

	userModel := reqData.ToModel()
	createdUser, err := ctrl.userService.CreateUser(userModel)
	if err != nil {
		return response.ErrorResponse(c, fiber.StatusInternalServerError, err, "Failed to create user")
	}

	return response.SuccessDataResponse(c, fiber.StatusCreated, dto.ToUserResponse(createdUser), "User Created Successfully")
}

func (ctrl *userController) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := ctrl.userService.GetUserById(id)
	if err != nil {
		return response.ErrorResponse(c, fiber.StatusInternalServerError, err, "Failed to fetch user")
	}

	if user == nil {
		return response.ErrorResponse(c, fiber.StatusNotFound, fmt.Errorf("user not found"), "User not found")
	}

	return response.SuccessDataResponse(c, fiber.StatusOK, dto.ToUserResponse(user), "User fetched successfully")
}
