package utils

import (
	"eleva/src/utils/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendEmailSuccess(t *testing.T) {
	assert := assert.New(t)

	t.Setenv("EMAIL_SENDER", "0680e6c3ef4c4d@example.com")
	t.Setenv("EMAIL_USERNAME", "0680e6c3ef4c4d")
	t.Setenv("EMAIL_SENDER_PASSWORD", "965685c0a2ab65")
	t.Setenv("EMAIL_HOST", "sandbox.smtp.mailtrap.io")
	t.Setenv("EMAIL_PORT", "587")

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

	email := Email{
		To:      tests.GenerateRandomEmail(),
		Subject: "random email",
		Body:    "this is a test email",
	}
	err := email.Send()
	assert.Error(err)
}
