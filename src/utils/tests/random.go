package tests

import (
	"crypto/rand"
	"io"
)

func GenerateRandomName(length int) string {
	characters := []byte("abcdefghijklmopqrstuvwxyz")

	name := make([]byte, length)
	io.ReadAtLeast(rand.Reader, name, length)

	for i := 0; i < length; i++ {
		name[i] = characters[int(name[i])%len(characters)]
	}

	return string(name)
}

func GenerateRandomEmail() string {
	return GenerateRandomName(10) + "@" + GenerateRandomName(5) + ".com"
}

func GenerateRandomPassword(length int) string {
	characters := []byte("abcdefghijklmopqrstuvwxyz123456789!@#$%&*()-_=+")

	password := make([]byte, length)
	io.ReadAtLeast(rand.Reader, password, length)

	for i := 0; i < length; i++ {
		password[i] = characters[int(password[i])%len(characters)]
	}

	return string(password)
}
