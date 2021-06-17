package main

import (
	"context"
	rooms "hciengserver/src/apps/chat/services"
	"hciengserver/src/database"
	"hciengserver/src/hciengserver"
	"hciengserver/src/helpers"
	"hciengserver/src/router"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	helpers.LoadEnv()
	hciengserver.InitSettings()
	database.Connect()
	spinUpRooms()
	r := router.InitRouter()
	r.Run(":" + os.Getenv("PORT"))
}

func spinUpRooms() {
	var rooms []*rooms.Room
	cursor, _ := database.GetMongoDBConn().Client().Database(hciengserver.DB_NAME).Collection(hciengserver.ROOMS_COLL).Find(context.Background(), bson.D{{}})
	cursor.All(context.Background(), &rooms)
	for _, room := range rooms {
		room.InitAndServe()
	}
}
