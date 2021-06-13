package auth

import (
	"hciengserver/src/apps/auth/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterEmailRoutes(router *gin.Engine) {
	api := router.Group("/auth")
	api.POST("login", controllers.Login)
	api.POST("register", controllers.Register)
}
