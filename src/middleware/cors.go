package middleware

import (
	"fmt"
	"hciengserver/src/hciengserver"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	fmt.Println("ok")
	return cors.New(cors.Config{
		AllowOrigins:     []string{hciengserver.DOMAIN},
		AllowMethods:     []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowHeaders:     []string{"X-Requested-With", "Content-Type"},
	})
}
