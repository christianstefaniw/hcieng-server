package email

import "github.com/gin-gonic/gin"

func RegisterEmailRoutes(router *gin.Engine) {
	api := router.Group("/email")

	api.POST("", email)
}
