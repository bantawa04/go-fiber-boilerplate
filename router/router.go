package router

import (
	"github.com/bantawao4/gofiber-boilerplate/router/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type Router struct {
	app          *fiber.App
	healthRouter *api.HealthRouter
	userRouter   *api.UserRouter
	todoRouter   *api.TodoRouter
}

func New(app *fiber.App) *Router {
	return &Router{
		app:          app,
		healthRouter: api.NewHealthRouter(app),
		userRouter:   api.NewUserRouter(app),
		todoRouter:   api.NewTodoRouter(app),
	}
}

func Setup(app *fiber.App) {
	router := New(app)
	app.Stack()
	// Setup API routes with rate limiter
	api := app.Group("/api", limiter.New())

	// Setup individual route groups
	router.healthRouter.Setup(api)
	router.userRouter.Setup(api)
	router.todoRouter.Setup(api)
}
