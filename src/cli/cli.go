package main

import (
	"flag"
	"fmt"
	accounts "hciengserver/src/apps/account/services"
	auth "hciengserver/src/apps/auth/services"
	"hciengserver/src/database"
	"hciengserver/src/hciengserver"
	"hciengserver/src/helpers"
	"log"
)

func init() {
	helpers.LoadEnv()
	hciengserver.InitSettings()
	database.Connect()
}

func createAccount(email, pass string, isAdmin bool) {
	newAccount := accounts.CreateAccount(email, pass, isAdmin)
	err := auth.ValidateAndAddAccountToDb(newAccount)
	if err != nil {
		log.Fatal("error adding account to db: ", err)
	}
	fmt.Printf("added account to db (email=%s, pass=%s, is_admin=%t)\n", email, pass, isAdmin)
}

func main() {
	createAccountPtr := flag.Bool("create-account", false, "create an account")
	emailPtr := flag.String("email", "test@gmail.com", "set account email")
	passPtr := flag.String("pass", "1234", "set account password")
	adminPtr := flag.Bool("admin", false, "make the account an admin account")

	flag.Parse()

	if createAccountPtr != nil {
		createAccount(*emailPtr, *passPtr, *adminPtr)
	}
}
