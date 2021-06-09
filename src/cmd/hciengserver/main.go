package main

import (
	"hciengserver/src/helpers"
	"hciengserver/src/router"
	"os"
)

func main() {
	helpers.LoadEnv()
	r := router.InitRouter()
	r.Run(":" + os.Getenv("PORT"))
}
