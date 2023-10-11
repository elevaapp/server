package users

import (
	"eleva/src/database"
	"eleva/src/database/models"
	"eleva/src/utils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CreateUserBody struct {
	Name     string `json:"name" validate:"required,alphanum"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Create(ctx *fiber.Ctx) error {
	body, errors := utils.ValidateBody[CreateUserBody](ctx)
	if len(errors) != 0 {
		return utils.ThrowError(ctx, 400, errors)
	}

	userId, err := utils.GenerateSnowflake()
	if err != nil {
		return utils.ThrowError(ctx, 502, err)
	}

	possibleUser := models.User{}

	database.Client.Where("name = ?", body.Name).Or("email = ?", body.Email).First(&possibleUser)
	if possibleUser.Name != "" {
		return utils.ThrowError(ctx, 402, "an user with this email or username already exists")
	}

	user := models.User{
		Id:       userId,
		Name:     body.Name,
		Email:    body.Email,
		Password: utils.EncryptPasswordToPBKDF2(body.Password),
	}
	user.Token = utils.GenerateToken(user)

	database.Client.Create(&user)

	verificationCode, err := utils.GenerateVerificationCode()
	if err != nil {
		return utils.ThrowError(ctx, 502, err)
	}

	database.Client.Create(&models.VerificationCode{
		UserId:    userId,
		Code:      verificationCode,
		ExpiresAt: time.Now().Unix() + 1800,
	})

	emailBody := fmt.Sprintf("You have just created a new account at the perfect workout organization app for you.\n Though, you still need to verify your account. Here's your verification code: %s. It will expire in 30 minutes.", verificationCode)

	email := utils.Email{
		To:      body.Email,
		Subject: "Your new account at Eleva",
		Body:    emailBody,
	}

	err = email.Send()
	if err != nil {
		return utils.ThrowError(ctx, 502, err)
	}

	return ctx.SendStatus(201)
}
