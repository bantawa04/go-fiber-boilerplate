package bootstrap

import (
	"log"
	"os"

	"github.com/bantawao4/gofiber-boilerplate/config"
	"github.com/bantawao4/gofiber-boilerplate/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewApplication() *fiber.App {

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
	})
	config.ConnectDb()

	app.Use(idempotency.New())

	app.Use(recover.New())

	// Log errors (status code >= 400) to error.log
	errorLogFile, err := os.OpenFile("./logs/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening error log file: %v", err)
	}
	// Remove the defer here as we want the file to stay open

	// Add general logger for all requests
	app.Use(logger.New())

	// Add specific error logger
	app.Use(logger.New(logger.Config{
		TimeZone: "UTC",
		Format:   "[${time}] ${status} ${latency} ${method} ${path}\n",
		Done: func(c *fiber.Ctx, logString []byte) {
			if c.Response().StatusCode() >= 400 {
				_, err := errorLogFile.Write(logString)
				if err != nil {
					log.Printf("failed to write to error log: %v", err)
				}
			}
		},
	}))

	router.Setup(app)

	return app
}
