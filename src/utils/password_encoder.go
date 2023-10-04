package utils

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"

	"golang.org/x/crypto/pbkdf2"
)

func EncryptPasswordToPBKDF2(password string) string {
	saltBytes := bytes.NewBufferString(GetEnv("SALT_VALUE")).Bytes()
	encryptedPassword := pbkdf2.Key([]byte(password), saltBytes, 4096, 32, sha1.New)

	encodedPassword := base64.StdEncoding.EncodeToString(encryptedPassword)
	return encodedPassword
}

func CompareEncryptedPasswords(target, original string) bool {
	saltBytes := bytes.NewBufferString(GetEnv("SALT_VALUE")).Bytes()
	encryptedTarget := pbkdf2.Key([]byte(target), saltBytes, 4096, 32, sha1.New)

	x, _ := base64.StdEncoding.DecodeString(original)
	difference := uint64(len(x)) ^ uint64(len(encryptedTarget))

	for i := 0; i < len(x) && i < len(encryptedTarget); i++ {
		difference |= uint64(x[i]) ^ uint64(encryptedTarget[i])
	}

	return difference == 0
}
