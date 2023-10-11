package users

import (
	"bytes"
	"eleva/src/database"
	"eleva/src/validation"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"eleva/src/utils/tests"

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

func TestCreateUserSuccess(t *testing.T) {
	assert := assert.New(t)
	app := fiber.New()
	validation.LoadValidator()

	t.Setenv("EMAIL_SENDER", "0680e6c3ef4c4d@example.com")
	t.Setenv("EMAIL_USERNAME", "0680e6c3ef4c4d")
	t.Setenv("EMAIL_SENDER_PASSWORD", "965685c0a2ab65")
	t.Setenv("EMAIL_HOST", "sandbox.smtp.mailtrap.io")
	t.Setenv("EMAIL_PORT", "587")

	app.Post("/users", Create)

	body := CreateUserBody{
		Name:     tests.GenerateRandomName(5),
		Email:    tests.GenerateRandomEmail(),
		Password: tests.GenerateRandomPassword(5),
	}
	requestBody, err := json.Marshal(body)
	assert.NoError(err)

	reader := bytes.NewReader(requestBody)
	request := httptest.NewRequest(http.MethodPost, "/users", reader)
	request.Header.Set("Content-Type", "application/json")

	response, _ := app.Test(request, -1)
	assert.Equal(fiber.StatusCreated, response.StatusCode)
}

func TestCreateUserInvalidName(t *testing.T) {
	assert := assert.New(t)
	app := fiber.New()
	validation.LoadValidator()

	app.Post("/users", Create)

	body := CreateUserBody{
		Name:     "",
		Email:    tests.GenerateRandomEmail(),
		Password: tests.GenerateRandomPassword(5),
	}
	requestBody, _ := json.Marshal(body)
	reader := bytes.NewReader(requestBody)
	request := httptest.NewRequest(http.MethodPost, "/users", reader)
	request.Header.Set("Content-Type", "application/json")

	response, _ := app.Test(request, -1)
	assert.Equal(fiber.StatusBadRequest, response.StatusCode)
}

func TestCreateUserInvalidEmail(t *testing.T) {
	assert := assert.New(t)
	app := fiber.New()
	validation.LoadValidator()

	app.Post("/users", Create)

	body := CreateUserBody{
		Name:     tests.GenerateRandomName(5),
		Email:    "email",
		Password: tests.GenerateRandomPassword(5),
	}
	requestBody, _ := json.Marshal(body)
	reader := bytes.NewReader(requestBody)
	request := httptest.NewRequest(http.MethodPost, "/users", reader)
	request.Header.Set("Content-Type", "application/json")

	response, _ := app.Test(request, -1)
	assert.Equal(fiber.StatusBadRequest, response.StatusCode)
}

func TestCreateUserInvalidPassword(t *testing.T) {
	assert := assert.New(t)
	app := fiber.New()
	validation.LoadValidator()

	app.Post("/users", Create)

	body := CreateUserBody{
		Name:     tests.GenerateRandomName(5),
		Email:    tests.GenerateRandomEmail(),
		Password: "",
	}
	requestBody, _ := json.Marshal(body)
	reader := bytes.NewReader(requestBody)
	request := httptest.NewRequest(http.MethodPost, "/users", reader)
	request.Header.Set("Content-Type", "application/json")

	response, _ := app.Test(request, -1)
	assert.Equal(fiber.StatusBadRequest, response.StatusCode)
}

func TestCreateUserInvalidBody(t *testing.T) {
	assert := assert.New(t)
	app := fiber.New()
	validation.LoadValidator()

	app.Post("/users", Create)

	request := httptest.NewRequest(http.MethodPost, "/users", nil)
	request.Header.Set("Content-Type", "application/json")

	response, _ := app.Test(request, -1)
	assert.Equal(fiber.StatusBadRequest, response.StatusCode)
}
