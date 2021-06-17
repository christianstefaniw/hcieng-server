package middleware

import (
	"errors"
	"hciengserver/src/apps/chat/services"
	"hciengserver/src/hciengserver"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterToMustJoinRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		user, ok := c.Get("user")
		if !ok {
			c.AbortWithError(http.StatusUnauthorized, errors.New("token not sent in request"))
			return
		}
		services.JoinRoom(hciengserver.MUST_JOIN_ROOM_ID, user)
	}
}
