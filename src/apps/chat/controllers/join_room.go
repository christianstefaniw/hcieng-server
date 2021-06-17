package controllers

import (
	"errors"
	"hciengserver/src/apps/chat/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JoinRoom(c *gin.Context) {
	rmId := c.Param("id")

	user, ok := c.Get("user")
	if !ok {
		c.AbortWithError(http.StatusUnauthorized, errors.New("token not sent in request"))
		return
	}

	services.JoinRoom(rmId, user)
}
