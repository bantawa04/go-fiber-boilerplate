package bootstrap

import (
	"github.com/bantawao4/gofiber-boilerplate/app/middleware"
	"github.com/bantawao4/gofiber-boilerplate/config"
	"github.com/bantawao4/gofiber-boilerplate/router"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewApplication() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
		ErrorHandler:  middleware.ErrorHandler,
	})
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
		CacheAge: 60,
	}

	app.Use(swagger.New(cfg))

	config.ConnectDb()

	app.Use(idempotency.New())

	app.Use(recover.New())

	app.Use(config.SetupLogger())

	router.Setup(app)

	return app
}
