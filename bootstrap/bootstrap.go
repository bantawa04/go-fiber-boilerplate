package bootstrap

import (
	"fmt"
	"log"
	"os"

	"github.com/bantawao4/gofiber-boilerplate/app/middleware"
	"github.com/bantawao4/gofiber-boilerplate/config"
	"github.com/bantawao4/gofiber-boilerplate/router"
	"github.com/gofiber/contrib/swagger"
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

	// Log errors (status code >= 400) to error.log
	errorLogFile, err := os.OpenFile("./logs/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening error log file: %v", err)
	}

	// Add general logger for all requests
	app.Use(logger.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
		Format:     "\n" + `{"timestamp":"${time}", "status":${status}, "latency":"${latency}", "method":"${method}", "path":"${path}", "error":"${error}", "response":[${resBody}]}`,
		Output:     errorLogFile,
		Done: func(c *fiber.Ctx, logString []byte) {
			if c.Response().StatusCode() >= 400 {
				fmt.Println("printing to error.log")
				errorLogFile.Write(logString)
				defer errorLogFile.Close()
			}
		},
	}))
	router.Setup(app)

	return app
}
