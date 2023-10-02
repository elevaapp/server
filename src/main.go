package main

import (
	"log"
	"the-perfect-workout-organizer/src/database"
	"the-perfect-workout-organizer/src/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.MountRoutes(app)

	database.Connect()

	log.Println(app.Listen(":3000"))
}
