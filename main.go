package main

import (
	"log"

	"github.com/bantawao4/gofiber-boilerplate/bootstrap"
)

func main() {
	app := bootstrap.NewApplication()
	log.Fatal(app.Listen(":8080"))
}
