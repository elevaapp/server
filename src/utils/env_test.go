package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvSuccess(t *testing.T) {
	assert := assert.New(t)

	t.Setenv("TEST_KEY", "value")

	testKey := GetEnv("TEST_KEY")
	assert.NotEmpty(testKey)
}

func TestGetEnvNotFound(t *testing.T) {
	assert := assert.New(t)

	testKey := GetEnv("TEST_KEY")
	assert.Empty(testKey)
}
