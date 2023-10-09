package models

type VerificationCode struct {
	UserId    string `json:"user_id" gorm:"primary_key"`
	Code      string `json:"code"`
	ExpiresAt int64  `json:"expires_at"`
}
