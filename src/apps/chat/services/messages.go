package services

import (
	accounts "hciengserver/src/apps/account/services"
	"time"
)

type message struct {
	Content string            `json:"message"`
	Time    time.Time         `json:"time"`
	Sender  *accounts.Account `json:"sender"`
}

func msgStringToMessage(msg []byte, sender *accounts.Account) *message {
	return &message{
		Content: string(msg),
		Sender:  sender,
		Time:    time.Now(),
	}
}
