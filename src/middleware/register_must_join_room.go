package middleware

import (
	"errors"
	"hciengserver/src/apps/chat/services"
	"hciengserver/src/hciengserver"
	"net/http"

	"github.com/gin-gonic/gin"
)

// everyone must join certain rooms, so this function auto joins new registers
func RegisterToMustJoinRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		user, ok := c.Get("user")
		if !ok {
			c.AbortWithError(http.StatusUnauthorized, errors.New("token not sent in request"))
			return
		}
		services.JoinRoom(hciengserver.HCI_ENG_ROOM_ID, user)
		services.JoinRoom(hciengserver.ANNOUNCEMENTS_ROOM_ID, user)
	}
}
