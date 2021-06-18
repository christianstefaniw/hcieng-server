package hciengserver

import (
	"os"
)

var (
	JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))
	DOMAIN     string
)

const (
	DEBUG                 = false
	PORT                  = "8080"
	EMAIL                 = "humbersideci.eng@gmail.com"
	DB_NAME               = "hciengonline"
	ACCOUNT_COLL          = "accounts"
	ROOMS_COLL            = "rooms"
	HCI_ENG_ROOM_ID       = "60cbe63c7b127cb2a206eeaf"
	ANNOUNCEMENTS_ROOM_ID = "60cbe61fb06d3f760cd6a4a6"
)

func InitSettings() {
	if DEBUG {
		DOMAIN = "http://localhost:3000"
	} else {
		DOMAIN = "www.hcieng.xyz"
	}
}
