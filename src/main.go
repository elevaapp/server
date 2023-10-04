package main

import (
	"eleva/src/database"
	"eleva/src/router"
	"eleva/src/validation"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.MountRoutes(app)

	database.Connect()
	validation.LoadValidator()

	log.Println(app.Listen(":3000"))
}
