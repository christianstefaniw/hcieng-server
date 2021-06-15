package account

import (
	account "hciengserver/src/apps/account/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func accountInfo(c *gin.Context) {
	c.JSON(http.StatusOK, c.Keys["user"].(*account.Account))
}
