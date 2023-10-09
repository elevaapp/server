package users

import (
	"eleva/src/database"
	"eleva/src/database/models"
	"eleva/src/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type VerifyUserByIdBody struct {
	Code string `json:"code" validate:"required,number"`
}

func VerifyById(ctx *fiber.Ctx) error {
	body, errors := utils.ValidateBody[VerifyUserByIdBody](ctx)
	if len(errors) != 0 {
		return utils.ThrowError(ctx, fiber.StatusBadRequest, errors)
	}

	fmt.Println(ctx.GetReqHeaders())

	userToken := ctx.GetReqHeaders()["Authorization"]
	if userToken == "" {
		return utils.ThrowError(ctx, fiber.StatusBadRequest, "invalid authorization identifier")
	}

	user := &models.User{}
	database.Client.Where("token = ?", userToken).First(user)

	if user.Id == "" {
		return utils.ThrowError(ctx, fiber.StatusNotFound, "user not found")
	}

	verificationCode := &models.VerificationCode{}
	database.Client.Where("code = ?", body.Code).Where("user_id = ?", user.Id).First(verificationCode)

	if verificationCode.ExpiresAt == 0 {
		return utils.ThrowError(ctx, fiber.StatusNotFound, "invalid verification code")
	}

	// Later there will be a goroutine that automatically deletes expired verification codes
	// so it is not needed to include this conditional

	database.Client.Model(user).Update("authorized", true)
	database.Client.Delete(verificationCode)

	return ctx.SendStatus(200)
}
