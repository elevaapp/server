package utils

import (
	"eleva/src/utils/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendEmailSuccess(t *testing.T) {
	assert := assert.New(t)

	t.Setenv("EMAIL_SENDER", GetEnv("FAKE_EMAIL_SENDER"))
	t.Setenv("EMAIL_USERNAME", GetEnv("FAKE_EMAIL_USERNAME"))
	t.Setenv("EMAIL_SENDER_PASSWORD", GetEnv("FAKE_EMAIL_SENDER_PASSWORD"))

	email := Email{
		To:      tests.GenerateRandomEmail(),
		Subject: "random email",
		Body:    "this is a test email",
	}
	err := email.Send()
	assert.NoError(err)
}

func TestSendEmailInvalidCredentials(t *testing.T) {
	assert := assert.New(t)

	t.Setenv("EMAIL_SENDER", "")
	t.Setenv("EMAIL_USERNAME", "")
	t.Setenv("EMAIL_SENDER_PASSWORD", "")

	email := Email{
		To:      tests.GenerateRandomEmail(),
		Subject: "random email",
		Body:    "this is a test email",
	}
	err := email.Send()
	assert.Error(err)
}
