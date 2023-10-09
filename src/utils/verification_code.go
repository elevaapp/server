package utils

import (
	"crypto/rand"
	"io"
)

func GenerateVerificationCode() (string, error) {
	numbers := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	code := make([]byte, 6)
	n, err := io.ReadAtLeast(rand.Reader, code, 6)
	if n != 6 {
		return "", err
	}

	for i := 0; i < len(code); i++ {
		code[i] = numbers[int(code[i])%len(numbers)]
	}

	return string(code), nil
}
