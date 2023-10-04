package users

import (
	"eleva/src/database"
	"eleva/src/database/models"
	"eleva/src/utils"

	"github.com/gofiber/fiber/v2"
)

type CreateUserBody struct {
	Name     string `json:"name" validate:"required,alphanum"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,alphanum"`
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

	email := utils.Email{
		To:      body.Email,
		Subject: "Your new account at Eleva",
		Body:    "You have just created a new account at the perfect workout organization app for you",
	}

	err = email.Send()
	if err != nil {
		return utils.ThrowError(ctx, 502, err)
	}

	return ctx.SendStatus(201)
}
