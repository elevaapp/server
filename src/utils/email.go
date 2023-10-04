package utils

import (
	"strconv"

	"github.com/go-mail/mail/v2"
)

type Email struct {
	To      string
	Subject string
	Body    string
}

func (email *Email) Send() error {
	sender := GetEnv("EMAIL_SENDER")
	emailPassword := GetEnv("EMAIL_SENDER_PASSWORD")
	host := GetEnv("EMAIL_HOST")
	port, _ := strconv.Atoi(GetEnv("EMAIL_PORT"))

	newMail := mail.NewMessage()

	newMail.SetHeaders(map[string][]string{
		"From":    {sender},
		"To":      {email.To},
		"Subject": {email.Subject},
	})
	newMail.SetBody("text/html", email.Body)

	dialer := mail.NewDialer(host, port, sender, emailPassword)
	return dialer.DialAndSend(newMail)
}
