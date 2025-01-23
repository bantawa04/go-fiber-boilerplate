package main

import (
	"github.com/bantawao4/gofiber-boilerplate/bootstrap"
	"log"
)

func main() {
	app := bootstrap.NewApplication()
	log.Fatal(app.Listen(":8080"))
}
