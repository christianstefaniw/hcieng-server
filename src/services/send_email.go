package services

import (
	"hciengserver/src/hciengserver"
	"net/smtp"
	"os"
)

func SendEmail(msg, to string) error {
	return smtp.SendMail("smtp.gmail.com:587", authEmail(), hciengserver.EMAIL, []string{to}, []byte(msg))
}

func authEmail() smtp.Auth {
	from := hciengserver.EMAIL
	pass := os.Getenv("EMAIL_PASS")

	return smtp.PlainAuth("", from, pass, "smtp.gmail.com")
}
