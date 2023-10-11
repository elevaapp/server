package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSnowflake(t *testing.T) {
	assert := assert.New(t)

	snowflake, err := GenerateSnowflake()
	assert.NoError(err)
	assert.NotEmpty(snowflake)
}
