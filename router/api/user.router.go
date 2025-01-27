package router

import (
	"fiber-boilerplate/app/controller"
	"fiber-boilerplate/app/repository"
	"fiber-boilerplate/app/service"
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
	users.Get("/", r.userController.GetUsers)
}