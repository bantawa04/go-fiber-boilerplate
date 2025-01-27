package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type Router struct {
	app          *fiber.App
	healthRouter *HealthRouter
	userRouter   *UserRouter
}

func New(app *fiber.App) *Router {
	return &Router{
		app:          app,
		healthRouter: NewHealthRouter(app),
		userRouter:   NewUserRouter(app),
	}
}

func Setup(app *fiber.App) {
	router := New(app)
	
	// Setup API routes with rate limiter
	api := app.Group("/api", limiter.New())
	
	// Setup individual route groups
	router.healthRouter.Setup(api)
	router.userRouter.Setup(api)
}