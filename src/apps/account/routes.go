package account

import (
	"hciengserver/src/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAccountRoutes(router *gin.Engine) {
	api := router.Group("/account")
	api.GET("info", middleware.WithAuth(), accountInfo)
}
