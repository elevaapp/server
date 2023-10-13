package utils

import (
	"fmt"
	"strconv"

	mail "github.com/xhit/go-simple-mail"
)

type Email struct {
	To      string
	Subject string
	Body    string
}

func (email *Email) Send() error {
	username := GetEnv("EMAIL_USERNAME")
	sender := GetEnv("EMAIL_SENDER")
	emailPassword := GetEnv("EMAIL_SENDER_PASSWORD")
	host := GetEnv("EMAIL_HOST")
	port, _ := strconv.Atoi(GetEnv("EMAIL_PORT"))

	server := mail.NewSMTPClient()
	server.Host = host
	server.Port = port
	server.Username = username
	server.Password = emailPassword
	server.Encryption = mail.EncryptionTLS

	client, err := server.Connect()
	if err != nil {
		return err
	}
	newMail := mail.NewMSG()

	newMail.SetFrom(sender).AddTo(email.To).SetSubject(email.Subject).SetBody(mail.TextHTML, email.Body)
	err = newMail.Send(client)
	if err != nil {
		return fmt.Errorf("could not send email: %s", err.Error())
	}

	return nil
}
