package account

import "github.com/gin-gonic/gin"

func accountInfo(c *gin.Context) {
	c.Writer.Write([]byte(c.Keys["user"].(string)))
}
