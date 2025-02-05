package api

import (
	"github.com/bantawao4/gofiber-boilerplate/app/controller"
	"github.com/bantawao4/gofiber-boilerplate/app/middleware"
	"github.com/bantawao4/gofiber-boilerplate/app/repository"
	"github.com/bantawao4/gofiber-boilerplate/app/service"
	"github.com/gofiber/fiber/v2"
)

type TodoRouter struct {
	app            *fiber.App
	todoController controller.TodoController
}

func NewTodoRouter(app *fiber.App) *TodoRouter {
	todoRepo := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepo)
	todoController := controller.NewTodoController(todoService)

	return &TodoRouter{
		app:            app,
		todoController: todoController,
	}
}

func (r *TodoRouter) Setup(api fiber.Router) {
	todos := api.Group("/todos")
	todos.Get("", r.todoController.GetTodos)
	todos.Get("/:id", r.todoController.GetTodoByID)
	todos.Post("", middleware.DBTransactionHandler(), r.todoController.CreateTodo)
	todos.Put("/:id", middleware.DBTransactionHandler(), r.todoController.UpdateTodo)
	todos.Delete("/:id", middleware.DBTransactionHandler(), r.todoController.DeleteTodo)
}
