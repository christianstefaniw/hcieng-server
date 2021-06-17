package auth

import (
	"hciengserver/src/apps/auth/oauth"
	standard "hciengserver/src/apps/auth/standard/controllers"
	"hciengserver/src/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	api := router.Group("/auth")
	api.POST("login", standard.Login)
	api.POST("login/google", oauth.GoogleAuthLogin)
	api.POST("register", middleware.RegisterToMustJoinRoom(), standard.Register)
	api.POST("register/google", middleware.RegisterToMustJoinRoom(), oauth.GoogleAuthRegister)
}
