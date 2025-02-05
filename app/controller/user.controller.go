package controller

import (
	"strconv"

	"github.com/bantawao4/gofiber-boilerplate/app/dto"
	"github.com/bantawao4/gofiber-boilerplate/app/middleware"
	"github.com/bantawao4/gofiber-boilerplate/app/request"
	"github.com/bantawao4/gofiber-boilerplate/app/response"
	"github.com/bantawao4/gofiber-boilerplate/app/service"
	"github.com/bantawao4/gofiber-boilerplate/app/validator"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetUsers(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
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

// GetUsers godoc
// @Summary Get list of users
// @Description Get paginated list of users with optional search
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param perPage query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.PaginationResponse{data=[]dto.UserResponse} "Users fetched successfully"
// @Router /users [get]
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

// CreateUser godoc
// @Summary Create new user
// @Description Create new user with the provided data
// @Tags users
// @Accept json
// @Produce json
// @Param user body request.CreateUserRequestData true "User data"
// @Success 201 {object} response.SuccessData{data=dto.UserResponse} "User created successfully"
// @Failure 400 {object} response.ErrorData "Bad request"
// @Failure 422 {object} response.Response{errors=[]response.ValidationError} "Validation error"
// @Router /users [post]
func (ctrl *userController) CreateUser(c *fiber.Ctx) error {
	// Get transaction from context
	tx := c.Locals(middleware.DBTransaction).(*gorm.DB)

	reqData := new(request.CreateUserRequestData)
	if err := c.BodyParser(reqData); err != nil {
		return err
	}

	if errors := ctrl.validator.Validate.Struct(reqData); errors != nil {
		return response.ValidationErrorResponse(c,
			ctrl.validator.GenerateValidationResponse(errors))
	}

	userModel := reqData.ToModel()

	createdUser, err := ctrl.userService.WithTrx(tx).CreateUser(userModel)
	if err != nil {
		return err
	}

	return response.SuccessDataResponse(c, fiber.StatusCreated,
		dto.ToUserResponse(createdUser), "User Created Successfully")
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get user details by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.SuccessData{data=dto.UserResponse} "User fetched successfully"
// @Failure 400 {object} response.ErrorData "Bad request"
// @Router /users/{id} [get]
func (ctrl *userController) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := ctrl.userService.GetUserById(id)
	if err != nil {
		return err
	}

	return response.SuccessDataResponse(c, fiber.StatusOK,
		dto.ToUserResponse(user), "User fetched successfully")
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user details by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body request.UpdateUserRequestData true "User data"
// @Success 200 {object}  response.SuccessData{data=dto.UserResponse} "User updated successfully"
// @Failure 400 {object} response.ErrorData "Bad request"
// @Failure 422 {object} response.ValidationError "Validation error"
// @Router /users/{id} [put]
func (ctrl *userController) UpdateUser(c *fiber.Ctx) error {
	tx := c.Locals(middleware.DBTransaction).(*gorm.DB)
	id := c.Params("id")

	reqData := new(request.UpdateUserRequestData)
	if err := c.BodyParser(reqData); err != nil {
		return err
	}

	if errors := ctrl.validator.Validate.Struct(reqData); errors != nil {
		return response.ValidationErrorResponse(c,
			ctrl.validator.GenerateValidationResponse(errors))
	}

	updatedUser, err := ctrl.userService.WithTrx(tx).UpdateUser(id, reqData.ToModel())
	if err != nil {
		return err
	}

	return response.SuccessDataResponse(c, fiber.StatusOK,
		dto.ToUserResponse(updatedUser), "User updated successfully")
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.Success "User deleted successfully"
// @Failure 400 {object} response.ErrorData "Bad request"
// @Router /users/{id} [delete]
func (ctrl *userController) DeleteUser(c *fiber.Ctx) error {
	tx := c.Locals(middleware.DBTransaction).(*gorm.DB)
	id := c.Params("id")

	err := ctrl.userService.WithTrx(tx).DeleteUser(id)
	if err != nil {
		return err
	}

	return response.SuccessResponse(c, fiber.StatusOK, "User deleted successfully")
}
