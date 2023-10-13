package users

import (
	"bytes"
	"eleva/src/api/auth"
	"eleva/src/database"
	"eleva/src/database/models"
	"eleva/src/utils"
	tests "eleva/src/utils/tests"
	"eleva/src/validation"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestVerifyUserSuccess(t *testing.T) {
	assert := assert.New(t)
	app := fiber.New()
	validation.LoadValidator()

	t.Setenv("EMAIL_SENDER", utils.GetEnv("FAKE_EMAIL_SENDER"))
	t.Setenv("EMAIL_USERNAME", utils.GetEnv("FAKE_EMAIL_USERNAME"))
	t.Setenv("EMAIL_SENDER_PASSWORD", utils.GetEnv("FAKE_EMAIL_SENDER_PASSWORD"))

	app.Post("/users", Create)
	app.Post("/auth/login", auth.Login)
	app.Post("/users/verify", Verify)

	name := tests.GenerateRandomName(5)
	email := tests.GenerateRandomEmail()
	password := tests.GenerateRandomPassword(5)

	createUserBody := CreateUserBody{
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

	loginBody := auth.LoginRequestBody{
		Name:     name,
		Password: password,
	}
	loginRequestBody, _ := json.Marshal(loginBody)
	loginReader := bytes.NewReader(loginRequestBody)
	loginRequest := httptest.NewRequest(http.MethodPost, "/auth/login", loginReader)
	loginRequest.Header.Set("Content-Type", "application/json")

	loginResponse, _ := app.Test(loginRequest, -1)
	assert.Equal(fiber.StatusOK, loginResponse.StatusCode)

	rawLoginResponseBody, err := io.ReadAll(loginResponse.Body)
	assert.NoError(err)

	var loginResponseBody auth.LoginResponseBody
	err = json.Unmarshal(rawLoginResponseBody, &loginResponseBody)
	assert.NoError(err)

	var user *models.User
	database.Client.Where("token = ?", loginResponseBody.Token).Find(&user)

	var verification *models.VerificationCode
	database.Client.Where("user_id = ?", user.Id).Find(&verification)

	verifyBody := VerifyUserBody{
		Code: verification.Code,
	}
	verifyRequestBody, _ := json.Marshal(verifyBody)
	verifyReader := bytes.NewReader(verifyRequestBody)
	verifyRequest := httptest.NewRequest(http.MethodPost, "/users/verify", verifyReader)
	verifyRequest.Header.Set("Content-Type", "application/json")
	verifyRequest.Header.Set("Authorization", loginResponseBody.Token)

	verifyResponse, _ := app.Test(verifyRequest, -1)
	assert.Equal(fiber.StatusOK, verifyResponse.StatusCode)
}

func TestVerifyUserMissingToken(t *testing.T) {
	assert := assert.New(t)
	app := fiber.New()
	validation.LoadValidator()

	t.Setenv("EMAIL_SENDER", utils.GetEnv("FAKE_EMAIL_SENDER"))
	t.Setenv("EMAIL_USERNAME", utils.GetEnv("FAKE_EMAIL_USERNAME"))
	t.Setenv("EMAIL_SENDER_PASSWORD", utils.GetEnv("FAKE_EMAIL_SENDER_PASSWORD"))

	app.Post("/users", Create)
	app.Post("/auth/login", auth.Login)
	app.Post("/users/verify", Verify)

	name := tests.GenerateRandomName(5)
	email := tests.GenerateRandomEmail()
	password := tests.GenerateRandomPassword(5)

	createUserBody := CreateUserBody{
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

	loginBody := auth.LoginRequestBody{
		Name:     name,
		Password: password,
	}
	loginRequestBody, _ := json.Marshal(loginBody)
	loginReader := bytes.NewReader(loginRequestBody)
	loginRequest := httptest.NewRequest(http.MethodPost, "/auth/login", loginReader)
	loginRequest.Header.Set("Content-Type", "application/json")

	loginResponse, _ := app.Test(loginRequest, -1)
	assert.Equal(fiber.StatusOK, loginResponse.StatusCode)

	rawLoginResponseBody, err := io.ReadAll(loginResponse.Body)
	assert.NoError(err)

	var loginResponseBody auth.LoginResponseBody
	err = json.Unmarshal(rawLoginResponseBody, &loginResponseBody)
	assert.NoError(err)

	var user *models.User
	database.Client.Where("token = ?", loginResponseBody.Token).Find(&user)

	var verification *models.VerificationCode
	database.Client.Where("user_id = ?", user.Id).Find(&verification)

	verifyBody := VerifyUserBody{
		Code: verification.Code,
	}
	verifyRequestBody, _ := json.Marshal(verifyBody)
	verifyReader := bytes.NewReader(verifyRequestBody)
	verifyRequest := httptest.NewRequest(http.MethodPost, "/users/verify", verifyReader)
	verifyRequest.Header.Set("Content-Type", "application/json")

	verifyResponse, _ := app.Test(verifyRequest, -1)
	assert.Equal(fiber.StatusUnauthorized, verifyResponse.StatusCode)
}

func TestVerifyUserInvalidToken(t *testing.T) {
	assert := assert.New(t)
	app := fiber.New()
	validation.LoadValidator()

	t.Setenv("EMAIL_SENDER", utils.GetEnv("FAKE_EMAIL_SENDER"))
	t.Setenv("EMAIL_USERNAME", utils.GetEnv("FAKE_EMAIL_USERNAME"))
	t.Setenv("EMAIL_SENDER_PASSWORD", utils.GetEnv("FAKE_EMAIL_SENDER_PASSWORD"))

	app.Post("/users", Create)
	app.Post("/auth/login", auth.Login)
	app.Post("/users/verify", Verify)

	name := tests.GenerateRandomName(5)
	email := tests.GenerateRandomEmail()
	password := tests.GenerateRandomPassword(5)

	createUserBody := CreateUserBody{
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

	loginBody := auth.LoginRequestBody{
		Name:     name,
		Password: password,
	}
	loginRequestBody, _ := json.Marshal(loginBody)
	loginReader := bytes.NewReader(loginRequestBody)
	loginRequest := httptest.NewRequest(http.MethodPost, "/auth/login", loginReader)
	loginRequest.Header.Set("Content-Type", "application/json")

	loginResponse, _ := app.Test(loginRequest, -1)
	assert.Equal(fiber.StatusOK, loginResponse.StatusCode)

	rawLoginResponseBody, err := io.ReadAll(loginResponse.Body)
	assert.NoError(err)

	var loginResponseBody auth.LoginResponseBody
	err = json.Unmarshal(rawLoginResponseBody, &loginResponseBody)
	assert.NoError(err)

	var user *models.User
	database.Client.Where("token = ?", loginResponseBody.Token).Find(&user)

	var verification *models.VerificationCode
	database.Client.Where("user_id = ?", user.Id).Find(&verification)

	verifyBody := VerifyUserBody{
		Code: verification.Code,
	}
	verifyRequestBody, _ := json.Marshal(verifyBody)
	verifyReader := bytes.NewReader(verifyRequestBody)
	verifyRequest := httptest.NewRequest(http.MethodPost, "/users/verify", verifyReader)
	verifyRequest.Header.Set("Content-Type", "application/json")
	verifyRequest.Header.Set("Authorization", "invalid token")

	verifyResponse, _ := app.Test(verifyRequest, -1)
	assert.Equal(fiber.StatusNotFound, verifyResponse.StatusCode)
}

func TestVerifyUserInvalidCode(t *testing.T) {
	assert := assert.New(t)
	app := fiber.New()
	validation.LoadValidator()

	t.Setenv("EMAIL_SENDER", utils.GetEnv("FAKE_EMAIL_SENDER"))
	t.Setenv("EMAIL_USERNAME", utils.GetEnv("FAKE_EMAIL_USERNAME"))
	t.Setenv("EMAIL_SENDER_PASSWORD", utils.GetEnv("FAKE_EMAIL_SENDER_PASSWORD"))

	app.Post("/users", Create)
	app.Post("/auth/login", auth.Login)
	app.Post("/users/verify", Verify)

	name := tests.GenerateRandomName(5)
	email := tests.GenerateRandomEmail()
	password := tests.GenerateRandomPassword(5)

	createUserBody := CreateUserBody{
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

	loginBody := auth.LoginRequestBody{
		Name:     name,
		Password: password,
	}
	loginRequestBody, _ := json.Marshal(loginBody)
	loginReader := bytes.NewReader(loginRequestBody)
	loginRequest := httptest.NewRequest(http.MethodPost, "/auth/login", loginReader)
	loginRequest.Header.Set("Content-Type", "application/json")

	loginResponse, _ := app.Test(loginRequest, -1)
	assert.Equal(fiber.StatusOK, loginResponse.StatusCode)

	rawLoginResponseBody, err := io.ReadAll(loginResponse.Body)
	assert.NoError(err)

	var loginResponseBody auth.LoginResponseBody
	err = json.Unmarshal(rawLoginResponseBody, &loginResponseBody)
	assert.NoError(err)

	verifyBody := VerifyUserBody{
		Code: "invalid code",
	}
	verifyRequestBody, _ := json.Marshal(verifyBody)
	verifyReader := bytes.NewReader(verifyRequestBody)
	verifyRequest := httptest.NewRequest(http.MethodPost, "/users/verify", verifyReader)
	verifyRequest.Header.Set("Content-Type", "application/json")
	verifyRequest.Header.Set("Authorization", loginResponseBody.Token)

	verifyResponse, _ := app.Test(verifyRequest, -1)
	assert.Equal(fiber.StatusBadRequest, verifyResponse.StatusCode)
}
