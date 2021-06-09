package email

import (
	"fmt"
	"hciengserver/src/hciengserver"
	"hciengserver/src/services"
)

type emailData struct {
	Name         string
	EmailAddress string
	Message      string
}

func send(info emailData) error {
	body := fmt.Sprintf("Email: %s\nName: %s\nMessage: %s\n", info.EmailAddress, info.Name, info.Message)
	msg :=
		"From: " + info.EmailAddress + "\n" +
			"Subject: Email from " + info.Name + "\n" +
			"To: " + hciengserver.EMAIL + "\n\n" +
			body

	return services.SendEmail(msg, hciengserver.EMAIL)
}
