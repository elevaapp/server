package router

import (
	"eleva/src/api/users"

	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App) {
	router := fiber.New()

	// /users
	router.Post("/users", users.Create)

	app.Mount("/", router)
}
