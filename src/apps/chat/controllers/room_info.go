package controllers

import (
	"context"
	"errors"
	accounts "hciengserver/src/apps/account/services"
	"hciengserver/src/apps/chat/services"
	"hciengserver/src/database"
	"hciengserver/src/hciengserver"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
				FindOne(context.Background(), query, options.FindOne()).
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

	upperMsgBound, err := strconv.Atoi(c.Param("upper_msg_bound"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	lowerMsgBound, err := strconv.Atoi(c.Param("lower_msg_bound"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if upperMsgBound > len(rm.Messages) {
		upperMsgBound = len(rm.Messages)
	}
	if lowerMsgBound > len(rm.Messages) {
		c.AbortWithError(http.StatusBadRequest, errors.New("all messages have been loaded"))
		return
	}

	rm.Messages = rm.Messages[lowerMsgBound:upperMsgBound]
	services.ReverseMessageSlice(rm.Messages)

	c.JSON(http.StatusOK, rm)
}
