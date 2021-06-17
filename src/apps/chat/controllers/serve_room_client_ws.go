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

	rm, ok := services.GetRoom(rmId)
	if !ok {
		c.AbortWithError(http.StatusNotFound, errors.New("room not active"))
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	user, ok := c.Get("user")
	if !ok {
		c.AbortWithError(http.StatusUnauthorized, errors.New("token not sent in request"))
		return
	}

	services.ServeWs(rm, user.(*accounts.Account), conn)
}
