package utils

import "github.com/gofiber/fiber/v2"

type Error struct {
	Message any `json:"message"`
}

func ThrowError(ctx *fiber.Ctx, statusCode int, message any) error {
	return ctx.Status(statusCode).JSON(Error{
		Message: message,
	})
}
