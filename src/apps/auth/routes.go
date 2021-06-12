package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterEmailRoutes(router *gin.Engine) {
	api := router.Group("/auth")
	api.POST("login", login)
}
