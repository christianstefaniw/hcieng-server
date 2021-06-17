package controllers

import (
	"errors"
	accounts "hciengserver/src/apps/account/services"
	"hciengserver/src/apps/chat/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRoom(c *gin.Context) {
	type RoomName struct {
		Name string `json:"name"`
	}

	var roomName RoomName
	if err := c.BindJSON(&roomName); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	room, err := services.NewRoomAndStore(roomName.Name)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	go room.Serve()

	userInterface, ok := c.Get("user")
	if !ok {
		c.AbortWithError(http.StatusBadRequest, errors.New("unauthorized"))
		return
	}

	user := userInterface.(*accounts.Account)
	user.AddRoom(room.Id.Hex())

	c.JSON(http.StatusCreated, room)
}
