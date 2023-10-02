package utils

import "github.com/gofiber/fiber/v2"

type Error struct {
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func ThrowError(ctx *fiber.Ctx, statusCode int, message, details string) error {
	return ctx.Status(statusCode).JSON(Error{
		Message: message,
		Details: details,
	})
}
