package hciengserver

import "os"

var (
	JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))
	DOMAIN     string
)

const (
	DEBUG        = true
	PORT         = "8080"
	EMAIL        = "humbersideci.eng@gmail.com"
	DB_NAME      = "hciengonline"
	ACCOUNT_COLL = "accounts"
)

func InitSettings() {
	if DEBUG {
		DOMAIN = "http://localhost:3000"
	} else {
		DOMAIN = "www.hcieng.xyz"
	}
}
