package account

import (
	account "hciengserver/src/apps/account/services"

	"github.com/gin-gonic/gin"
)

func accountInfo(c *gin.Context) {
	c.Writer.Write([]byte(c.Keys["user"].(*account.Account).EmailAddr))
}
