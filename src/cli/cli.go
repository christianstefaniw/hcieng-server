package main

import (
	"flag"
	"fmt"
	accounts "hciengserver/src/apps/account/services"
	auth "hciengserver/src/apps/auth/standard/services"
	chat "hciengserver/src/apps/chat/services"
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

func createAccount(email, pass, first, last string, isAdmin bool) {
	newAccount := accounts.CreateAccount(email, pass, first, last, isAdmin)
	err := auth.AddNewRegisterToDb(newAccount)
	if err != nil {
		log.Fatal("error adding account to db: ", err)
	}
	chat.JoinRoom(hciengserver.HCI_ENG_ROOM_ID, newAccount)
	chat.JoinRoom(hciengserver.ANNOUNCEMENTS_ROOM_ID, newAccount)
	fmt.Printf("added account to db (email=%s, pass=%s, is_admin=%t)\n", email, pass, isAdmin)
}

func createMustJoinRoom(roomName string, adminTextOnly bool) {
	room, err := chat.NewRoomAndStore(roomName, adminTextOnly)
	if err != nil {
		log.Fatal("error creating room: ", err)
	}
	fmt.Printf("created room with id `%s`\n", room.Id)
}

func main() {
	createAccountPtr := flag.Bool("create-account", false, "create an account")
	createMustJoinRoomPtr := flag.Bool("create-must-room", false, "create a room that all users auto join")
	adminTextOnlyPtr := flag.Bool("admin-text-only", false, "create a room that only admin can message")
	roomNamePtr := flag.String("name", "HCI Eng", "room name")
	emailPtr := flag.String("email", "test@gmail.com", "set account email")
	firstPtr := flag.String("first", "firstName", "first name")
	lastPtr := flag.String("last", "lastName", "last name")
	passPtr := flag.String("pass", "1234", "set account password")
	adminPtr := flag.Bool("admin", false, "make the account an admin account")

	flag.Parse()

	if *createAccountPtr {
		createAccount(*emailPtr, *passPtr, *firstPtr, *lastPtr, *adminPtr)
	}

	if *createMustJoinRoomPtr {
		createMustJoinRoom(*roomNamePtr, *adminTextOnlyPtr)
	}
}
