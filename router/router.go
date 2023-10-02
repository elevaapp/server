package router

import (
	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App) {
	router := fiber.New()

	// Insert routes

	app.Mount("/", router)
}
