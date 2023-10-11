package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateVerificationCodeSuccess(t *testing.T) {
	assert := assert.New(t)

	code, err := GenerateVerificationCode()
	assert.NoError(err)
	assert.NotEqual(code, "")
	assert.Len(code, 6)
}
