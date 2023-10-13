package utils

import (
	"bytes"
	"eleva/src/utils/tests"
	"eleva/src/validation"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestValidateBodySuccess(t *testing.T) {
	assert := assert.New(t)
	validation.LoadValidator()

	type BodySample struct {
		Name string `json:"name" validate:"required,alphanum"`
	}

	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {
		_, err := ValidateBody[BodySample](c)

		if len(err) != 0 {
			return ThrowError(c, fiber.StatusBadRequest, err)
		}

		return c.SendStatus(fiber.StatusOK)
	})

	body := BodySample{
		Name: tests.GenerateRandomName(5),
	}
	requestBody, err := json.Marshal(body)
	assert.NoError(err)

	requestReader := bytes.NewReader(requestBody)

	request := httptest.NewRequest(http.MethodPost, "/", requestReader)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request, -1)
	assert.NoError(err)
	assert.Equal(fiber.StatusOK, response.StatusCode)
}

func TestValidateBodyMissingBody(t *testing.T) {
	assert := assert.New(t)
	validation.LoadValidator()

	type BodySample struct {
		Name string `validate:"required,alphanum"`
	}

	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {
		_, err := ValidateBody[BodySample](c)

		if len(err) != 0 {
			return ThrowError(c, fiber.StatusBadRequest, err)
		}

		return c.SendStatus(fiber.StatusOK)
	})

	request := httptest.NewRequest(http.MethodPost, "/", nil)
	response, err := app.Test(request, -1)
	assert.NoError(err)
	assert.Equal(fiber.StatusBadRequest, response.StatusCode)
}

func TestValidateBodyInvalidField(t *testing.T) {
	assert := assert.New(t)
	validation.LoadValidator()

	type BodySample struct {
		Name string `validate:"required,alphanum"`
	}

	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {
		_, err := ValidateBody[BodySample](c)

		if len(err) != 0 {
			return ThrowError(c, fiber.StatusBadRequest, err)
		}

		return c.SendStatus(fiber.StatusOK)
	})

	body := BodySample{
		Name: "",
	}
	requestBody, err := json.Marshal(body)
	assert.NoError(err)

	requestReader := bytes.NewReader(requestBody)

	request := httptest.NewRequest(http.MethodPost, "/", requestReader)
	response, err := app.Test(request, -1)
	assert.NoError(err)
	assert.Equal(fiber.StatusBadRequest, response.StatusCode)
}
