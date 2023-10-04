package utils

import (
	"eleva/src/validation"

	"github.com/gofiber/fiber/v2"
)

func ValidateBody[T any](ctx *fiber.Ctx) (T, []string) {
	var body T
	err := ctx.BodyParser(&body)
	if err != nil {
		return body, []string{"invalid body"}
	}

	errors := validation.Validate(body)

	return body, errors
}
