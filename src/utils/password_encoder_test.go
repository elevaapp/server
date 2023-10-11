package utils

import (
	"eleva/src/utils/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptPasswordToPBKDF2(t *testing.T) {
	assert := assert.New(t)

	t.Setenv("SALT_VALUE", "565db310-6f51-436e-8be7-5ff91ed9a508") // random uuid

	password := tests.GenerateRandomPassword(5)
	encryptedPassword := EncryptPasswordToPBKDF2(password)
	assert.NotEmpty(encryptedPassword)
}

func TestCompareEncryptedPasswordsSuccess(t *testing.T) {
	assert := assert.New(t)

	t.Setenv("SALT_VALUE", "565db310-6f51-436e-8be7-5ff91ed9a508") // random uuid

	password := tests.GenerateRandomPassword(5)
	encryptedPassword := EncryptPasswordToPBKDF2(password)
	assert.True(CompareEncryptedPasswords(password, encryptedPassword))
}

func TestCompareEncryptedPasswordsUnequal(t *testing.T) {
	assert := assert.New(t)

	t.Setenv("SALT_VALUE", "565db310-6f51-436e-8be7-5ff91ed9a508") // random uuid

	password := tests.GenerateRandomPassword(5)
	encryptedPassword := EncryptPasswordToPBKDF2(password)
	assert.False(CompareEncryptedPasswords("123", encryptedPassword))
}
