package utils

import (
	"eleva/src/database/models"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/kcorlidy/dangerous"
)

func GenerateToken(user models.User) string {
	userIdEncodedToBase64 := base64.URLEncoding.EncodeToString([]byte(user.Id))
	timestampWhenTokenWasCreated := base64.URLEncoding.EncodeToString([]byte(fmt.Sprint(time.Now().Unix())))
	token := fmt.Sprintf("%s.%s.", userIdEncodedToBase64, timestampWhenTokenWasCreated)

	secret := GetEnv("ENCRYPTION_KEY")
	data := map[string]interface{}{
		"email":    user.Email,
		"password": user.Password,
	}

	serializer := dangerous.Serializer{
		Secret: secret,
		Salt:   "auth",
	}
	result, _ := serializer.URLSafeDumps(data)

	return token + string(result)
}
