package auth

import (
	"bytes"
	"eleva/src/api/users"
	"eleva/src/database"
	"eleva/src/utils/tests"
	"eleva/src/validation"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	database.MockConnection()

	exitCode := m.Run()

	instance, _ := database.Client.DB()
	instance.Close()

	os.Remove("./test.db")

	database.Client = nil

	os.Exit(exitCode)
}

func TestLoginSuccess(t *testing.T) {
	assert := assert.New(t)
	app := fiber.New()
	validation.LoadValidator()

	t.Setenv("EMAIL_SENDER", "0680e6c3ef4c4d@example.com")
	t.Setenv("EMAIL_USERNAME", "0680e6c3ef4c4d")
	t.Setenv("EMAIL_SENDER_PASSWORD", "965685c0a2ab65")
	t.Setenv("EMAIL_HOST", "sandbox.smtp.mailtrap.io")
	t.Setenv("EMAIL_PORT", "587")

	app.Post("/users", users.Create)
	app.Post("/auth/login", Login)

	name := tests.GenerateRandomName(5)
	email := tests.GenerateRandomEmail()
	password := tests.GenerateRandomPassword(5)

	createUserBody := users.CreateUserBody{
		Name:     name,
		Email:    email,
		Password: password,
	}
	createUserRequestBody, _ := json.Marshal(createUserBody)
	createUserReader := bytes.NewReader(createUserRequestBody)
	createUserRequest := httptest.NewRequest(http.MethodPost, "/users", createUserReader)
	createUserRequest.Header.Set("Content-Type", "application/json")

	createUserResponse, _ := app.Test(createUserRequest, -1)
	assert.Equal(fiber.StatusCreated, createUserResponse.StatusCode)

	loginBody := LoginRequestBody{
		Name:     name,
		Password: password,
	}
	loginRequestBody, _ := json.Marshal(loginBody)
	loginReader := bytes.NewReader(loginRequestBody)
	loginRequest := httptest.NewRequest(http.MethodPost, "/auth/login", loginReader)
	loginRequest.Header.Set("Content-Type", "application/json")

	loginResponse, _ := app.Test(loginRequest, -1)
	assert.Equal(fiber.StatusOK, loginResponse.StatusCode)
}

func TestLoginUserNotFound(t *testing.T) {
	assert := assert.New(t)
	app := fiber.New()
	validation.LoadValidator()

	t.Setenv("EMAIL_SENDER", "0680e6c3ef4c4d@example.com")
	t.Setenv("EMAIL_USERNAME", "0680e6c3ef4c4d")
	t.Setenv("EMAIL_SENDER_PASSWORD", "965685c0a2ab65")
	t.Setenv("EMAIL_HOST", "sandbox.smtp.mailtrap.io")
	t.Setenv("EMAIL_PORT", "587")

	app.Post("/users", users.Create)
	app.Post("/auth/login", Login)

	name := tests.GenerateRandomName(5)
	email := tests.GenerateRandomEmail()
	password := tests.GenerateRandomPassword(5)

	createUserBody := users.CreateUserBody{
		Name:     name,
		Email:    email,
		Password: password,
	}
	createUserRequestBody, _ := json.Marshal(createUserBody)
	createUserReader := bytes.NewReader(createUserRequestBody)
	createUserRequest := httptest.NewRequest(http.MethodPost, "/users", createUserReader)
	createUserRequest.Header.Set("Content-Type", "application/json")

	createUserResponse, _ := app.Test(createUserRequest, -1)
	assert.Equal(fiber.StatusCreated, createUserResponse.StatusCode)

	loginBody := LoginRequestBody{
		Name:     "a",
		Password: password,
	}
	loginRequestBody, _ := json.Marshal(loginBody)
	loginReader := bytes.NewReader(loginRequestBody)
	loginRequest := httptest.NewRequest(http.MethodPost, "/auth/login", loginReader)
	loginRequest.Header.Set("Content-Type", "application/json")

	loginResponse, _ := app.Test(loginRequest, -1)
	assert.Equal(fiber.StatusNotFound, loginResponse.StatusCode)
}

func TestLoginWrongPassword(t *testing.T) {
	assert := assert.New(t)
	app := fiber.New()
	validation.LoadValidator()

	t.Setenv("EMAIL_SENDER", "0680e6c3ef4c4d@example.com")
	t.Setenv("EMAIL_USERNAME", "0680e6c3ef4c4d")
	t.Setenv("EMAIL_SENDER_PASSWORD", "965685c0a2ab65")
	t.Setenv("EMAIL_HOST", "sandbox.smtp.mailtrap.io")
	t.Setenv("EMAIL_PORT", "587")

	app.Post("/users", users.Create)
	app.Post("/auth/login", Login)

	name := tests.GenerateRandomName(5)
	email := tests.GenerateRandomEmail()
	password := tests.GenerateRandomPassword(5)

	createUserBody := users.CreateUserBody{
		Name:     name,
		Email:    email,
		Password: password,
	}
	createUserRequestBody, _ := json.Marshal(createUserBody)
	createUserReader := bytes.NewReader(createUserRequestBody)
	createUserRequest := httptest.NewRequest(http.MethodPost, "/users", createUserReader)
	createUserRequest.Header.Set("Content-Type", "application/json")

	createUserResponse, _ := app.Test(createUserRequest, -1)
	assert.Equal(fiber.StatusCreated, createUserResponse.StatusCode)

	loginBody := LoginRequestBody{
		Name:     name,
		Password: "password",
	}
	loginRequestBody, _ := json.Marshal(loginBody)
	loginReader := bytes.NewReader(loginRequestBody)
	loginRequest := httptest.NewRequest(http.MethodPost, "/auth/login", loginReader)
	loginRequest.Header.Set("Content-Type", "application/json")

	loginResponse, _ := app.Test(loginRequest, -1)
	assert.Equal(fiber.StatusUnauthorized, loginResponse.StatusCode)
}
