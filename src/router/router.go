package router

import (
	"eleva/src/api/auth"
	"eleva/src/api/users"

	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App) {
	router := fiber.New()

	// auth
	router.Post("/auth/login", auth.Login)

	// /users
	router.Post("/users", users.Create)
	router.Post("/users/verify", users.Verify)

	app.Mount("/", router)
}
