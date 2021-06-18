package controllers

import (
	"errors"
	accounts "hciengserver/src/apps/account/services"
	"hciengserver/src/apps/chat/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func ServeRoom(c *gin.Context) {
	rmId := c.Param("id")

	user, ok := c.Get("user")
	if !ok {
		c.AbortWithError(http.StatusUnauthorized, errors.New("token not sent in request"))
		return
	}

	rm := services.GetRoom(rmId)
	if rm == nil {
		c.AbortWithError(http.StatusNotFound, errors.New("room not active"))
		return
	}

	if rm.CheckClientInRoom(user.(*accounts.Account).EmailAddr) {
		c.AbortWithStatus(http.StatusAlreadyReported)
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	services.ServeWs(rm, user.(*accounts.Account), conn)
}
