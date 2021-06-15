package router

import (
	"hciengserver/src/apps/account"
	"hciengserver/src/apps/auth"
	"hciengserver/src/apps/email"
	"hciengserver/src/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(middleware.Cors())

	email.RegisterEmailRoutes(router)
	auth.RegisterAuthRoutes(router)
	account.RegisterAccountRoutes(router)

	return router
}
