package main

import (
	"log"
	"the-end-store-api/database"
	"the-end-store-api/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.MountRoutes(app)

	database.Connect()

	log.Println(app.Listen(":3000"))
}
