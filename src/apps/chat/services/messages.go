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

func ReverseMessageSlice(s []*message) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
