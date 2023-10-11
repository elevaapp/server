package auth

import (
	"eleva/src/database"
	"eleva/src/database/models"
	"eleva/src/utils"

	"github.com/gofiber/fiber/v2"
)

type LoginRequestBody struct {
	Name     string `json:"name" validate:"required,alphanum"`
	Password string `json:"password" validate:"required"`
}

type LoginResponseBody struct {
	Token string `json:"token"`
}

func Login(ctx *fiber.Ctx) error {
	body, errs := utils.ValidateBody[LoginRequestBody](ctx)
	if len(errs) > 0 {
		return utils.ThrowError(ctx, fiber.StatusBadRequest, errs)
	}

	user := &models.User{}
	database.Client.Where("name = ?", body.Name).First(user)

	if user.Id == "" {
		return utils.ThrowError(ctx, fiber.StatusNotFound, "user not found")
	}

	isPasswordCorrect := utils.CompareEncryptedPasswords(body.Password, user.Password)

	if !isPasswordCorrect {
		return utils.ThrowError(ctx, fiber.StatusUnauthorized, "name or password are incorrect")
	}

	return ctx.Status(fiber.StatusOK).JSON(LoginResponseBody{
		Token: user.Token,
	})
}
