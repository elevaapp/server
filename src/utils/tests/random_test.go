package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomName(t *testing.T) {
	assert := assert.New(t)

	name := GenerateRandomName(5)
	assert.NotEmpty(name)
	assert.Len(name, 5)
}

func TestGenerateRandomEmail(t *testing.T) {
	assert := assert.New(t)

	email := GenerateRandomEmail()
	assert.NotEmpty(email)
}

func TestGenerateRandomPassword(t *testing.T) {
	assert := assert.New(t)

	password := GenerateRandomName(5)
	assert.NotEmpty(password)
	assert.Len(password, 5)
}
