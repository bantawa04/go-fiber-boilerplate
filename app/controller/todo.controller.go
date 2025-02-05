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

type TodoController interface {
	GetTodos(c *fiber.Ctx) error
	CreateTodo(c *fiber.Ctx) error
	GetTodoByID(c *fiber.Ctx) error
	UpdateTodo(c *fiber.Ctx) error
	DeleteTodo(c *fiber.Ctx) error
}

type todoController struct {
	todoService service.TodoService
	validator   validator.TodoValidator
}

func NewTodoController(todoService service.TodoService) TodoController {
	return &todoController{
		todoService: todoService,
		validator:   validator.NewTodoValidator(),
	}
}

// GetTodos godoc
// @Summary Get list of todos
// @Description Get paginated list of todos with optional search
// @Tags todos
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param perPage query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.PaginationResponse{data=[]dto.TodoResponse} "Todos fetched successfully"
// @Router /todos [get]
func (ctrl *todoController) GetTodos(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	perPage, err := strconv.Atoi(c.Query("perPage", "10"))
	if err != nil || perPage < 1 {
		perPage = 10
	}

	searchQuery := c.Query("search", "")

	todos, meta, err := ctrl.todoService.GetTodos(page, perPage, searchQuery)
	if err != nil {
		return err
	}

	return response.SuccessPaginationResponse(c, fiber.StatusOK, dto.ToTodoListResponse(todos), meta, "Todos fetched successfully")
}

// CreateTodo godoc
// @Summary Create new todo
// @Description Create new todo with the provided data
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body request.CreateTodoRequest true "Todo data"
// @Success 201 {object} response.SuccessData{data=dto.TodoResponse} "Todo created successfully"
// @Failure 400 {object} response.ErrorData "Bad request"
// @Failure 422 {object} response.Response{errors=[]response.ValidationError} "Validation error"
// @Router /todos [post]
func (ctrl *todoController) CreateTodo(c *fiber.Ctx) error {
	tx := c.Locals(middleware.DBTransaction).(*gorm.DB)

	reqData := new(request.CreateTodoRequest)
	if err := c.BodyParser(reqData); err != nil {
		return err
	}

	if errors := ctrl.validator.Validate.Struct(reqData); errors != nil {
		return response.ValidationErrorResponse(c,
			ctrl.validator.GenerateValidationResponse(errors))
	}

	todoModel := reqData.ToModel()

	createdTodo, err := ctrl.todoService.WithTrx(tx).CreateTodo(todoModel)
	if err != nil {
		return err
	}

	return response.SuccessDataResponse(c, fiber.StatusCreated,
		dto.ToTodoResponse(createdTodo), "Todo Created Successfully")
}

// GetTodoByID godoc
// @Summary Get todo by ID
// @Description Get todo details by todo ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} response.SuccessData{data=dto.TodoResponse} "Todo fetched successfully"
// @Failure 400 {object} response.ErrorData "Bad request"
// @Router /todos/{id} [get]
func (ctrl *todoController) GetTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")

	todo, err := ctrl.todoService.GetTodoById(id)
	if err != nil {
		return err
	}

	return response.SuccessDataResponse(c, fiber.StatusOK,
		dto.ToTodoResponse(todo), "Todo fetched successfully")
}

// UpdateTodo godoc
// @Summary Update todo
// @Description Update todo details by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param todo body request.CreateTodoRequest true "Todo data"
// @Success 200 {object} response.SuccessData{data=dto.TodoResponse} "Todo updated successfully"
// @Failure 400 {object} response.ErrorData "Bad request"
// @Failure 422 {object} response.ValidationError "Validation error"
// @Router /todos/{id} [put]
func (ctrl *todoController) UpdateTodo(c *fiber.Ctx) error {
	tx := c.Locals(middleware.DBTransaction).(*gorm.DB)
	id := c.Params("id")

	reqData := new(request.UpdateTodoRequest)
	if err := c.BodyParser(reqData); err != nil {
		return err
	}

	if errors := ctrl.validator.Validate.Struct(reqData); errors != nil {
		return response.ValidationErrorResponse(c,
			ctrl.validator.GenerateValidationResponse(errors))
	}

	updatedTodo, err := ctrl.todoService.WithTrx(tx).UpdateTodo(id, reqData.ToModel())
	if err != nil {
		return err
	}

	return response.SuccessDataResponse(c, fiber.StatusOK,
		dto.ToTodoResponse(updatedTodo), "Todo updated successfully")
}

// DeleteTodo godoc
// @Summary Delete todo
// @Description Delete todo by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} response.Success "Todo deleted successfully"
// @Failure 400 {object} response.ErrorData "Bad request"
// @Router /todos/{id} [delete]
func (ctrl *todoController) DeleteTodo(c *fiber.Ctx) error {
	tx := c.Locals(middleware.DBTransaction).(*gorm.DB)
	id := c.Params("id")

	err := ctrl.todoService.WithTrx(tx).DeleteTodo(id)
	if err != nil {
		return err
	}

	return response.SuccessResponse(c, fiber.StatusOK, "Todo deleted successfully")
}
