package api // Changed from 'router' to 'api'

import (
	"github.com/bantawao4/gofiber-boilerplate/app/controller"
	"github.com/bantawao4/gofiber-boilerplate/app/repository"
	"github.com/bantawao4/gofiber-boilerplate/app/service"
	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
	app            *fiber.App
	userController controller.UserController
}

func NewUserRouter(app *fiber.App) *UserRouter {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	return &UserRouter{
		app:            app,
		userController: userController,
	}
}

func (r *UserRouter) Setup(api fiber.Router) {
	users := api.Group("/users")
	users.Get("", r.userController.GetUsers)
	users.Post("", r.userController.CreateUser)
	users.Get("/:id", r.userController.GetUserByID)
}
