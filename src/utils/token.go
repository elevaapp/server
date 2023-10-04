package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"eleva/src/database/models"
	"encoding/base64"
	"fmt"
	"time"
)

func GenerateToken(user models.User) string {
	userIdEncodedToBase64 := base64.URLEncoding.EncodeToString([]byte(user.Id))
	timestampWhenTokenWasCreated := base64.URLEncoding.EncodeToString([]byte(fmt.Sprint(time.Now().Unix())))
	token := fmt.Sprintf("%s.%s.", userIdEncodedToBase64, timestampWhenTokenWasCreated)

	hmacComponentWithEmailAndPassword := hmac.New(sha256.New, []byte(fmt.Sprintf("%s+%s", user.Email, user.Password)))

	sha := base64.URLEncoding.EncodeToString(hmacComponentWithEmailAndPassword.Sum(nil))

	return token + sha
}
