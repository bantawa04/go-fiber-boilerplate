package main

import (
	"log"

	"github.com/bantawao4/gofiber-boilerplate/app/middleware"
	"github.com/bantawao4/gofiber-boilerplate/bootstrap"
)

func main() {
	app := bootstrap.NewApplication()
	app.Use(middleware.ErrorHandler)
	log.Fatal(app.Listen(":8080"))
}
