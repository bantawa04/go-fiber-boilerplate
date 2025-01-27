package router

import (
	apirouter "github.com/bantawao4/gofiber-boilerplate/router/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type Router struct {
	app          *fiber.App
	healthRouter *apirouter.HealthRouter
	userRouter   *apirouter.UserRouter
}

func New(app *fiber.App) *Router {
	return &Router{
		app:          app,
		healthRouter: apirouter.NewHealthRouter(app),
		userRouter:   apirouter.NewUserRouter(app),
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
}
