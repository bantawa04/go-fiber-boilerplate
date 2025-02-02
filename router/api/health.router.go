package api

import "github.com/gofiber/fiber/v2"

type HealthRouter struct {
	app *fiber.App
}

func NewHealthRouter(app *fiber.App) *HealthRouter {
	return &HealthRouter{
		app: app,
	}
}

func (r *HealthRouter) Setup(api fiber.Router) {
	api.Get("/", r.handleHealthCheck)
}

func (r *HealthRouter) handleHealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "API is running",
	})
}
