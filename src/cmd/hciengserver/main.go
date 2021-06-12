package main

import (
	"hciengserver/src/database"
	"hciengserver/src/hciengserver"
	"hciengserver/src/helpers"
	"hciengserver/src/router"
	"os"
)

func main() {
	helpers.LoadEnv()
	hciengserver.InitSettings()
	database.Connect()
	r := router.InitRouter()
	r.Run(":" + os.Getenv("PORT"))
}
