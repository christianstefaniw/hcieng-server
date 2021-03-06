package controllers

import (
	"context"
	"errors"
	accounts "hciengserver/src/apps/account/services"
	"hciengserver/src/apps/chat/services"
	"hciengserver/src/database"
	"hciengserver/src/hciengserver"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func AllUserRooms(c *gin.Context) {
	var allMinRooms []*services.MinRoomData

	user, ok := c.Get("user")
	if !ok {
		c.AbortWithError(http.StatusUnauthorized, errors.New("token not sent in request"))
		return
	}

	for _, rmId := range user.(*accounts.Account).Rooms {
		var rm services.MinRoomData
		query := bson.M{
			"_id": rmId,
		}
		database.GetMongoDBConn().Client().
			Database(hciengserver.DB_NAME).
			Collection(hciengserver.ROOMS_COLL).
			FindOne(context.Background(), query).Decode(&rm)
		allMinRooms = append(allMinRooms, &rm)
	}

	c.JSON(http.StatusOK, allMinRooms)
}
