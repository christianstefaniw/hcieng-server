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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RoomInfo(c *gin.Context) {
	var rm services.Room

	rmId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, ok := c.Get("user")
	if !ok {
		c.AbortWithError(http.StatusUnauthorized, errors.New("token not sent in request"))
		return
	}

	for _, userRoomId := range user.(*accounts.Account).Rooms {
		if userRoomId == rmId {
			query := bson.M{
				"_id": rmId,
			}
			err := database.GetMongoDBConn().Client().
				Database(hciengserver.DB_NAME).
				Collection(hciengserver.ROOMS_COLL).
				FindOne(context.Background(), query).
				Decode(&rm)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
		}
	}

	if rm.Id == (primitive.ObjectID{}) {
		c.AbortWithError(http.StatusNotFound, errors.New("you have not joined this room"))
		return
	}

	c.JSON(http.StatusOK, rm)
}
