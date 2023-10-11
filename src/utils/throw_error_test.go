package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestThrowErrorSuccess(t *testing.T) {
	assert := assert.New(t)

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return ThrowError(c, fiber.StatusBadRequest, "")
	})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response, err := app.Test(request, -1)
	assert.NoError(err)
	assert.Equal(fiber.StatusBadRequest, response.StatusCode)
}
