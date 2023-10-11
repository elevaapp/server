package utils

import (
	"eleva/src/database/models"
	"eleva/src/utils/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToke(t *testing.T) {
	assert := assert.New(t)

	snowflake, err := GenerateSnowflake()
	assert.NoError(err)

	user := models.User{
		Id:       snowflake,
		Name:     tests.GenerateRandomName(5),
		Email:    tests.GenerateRandomEmail(),
		Password: tests.GenerateRandomPassword(8),
	}

	token := GenerateToken(user)
	assert.NotEmpty(token)
}
