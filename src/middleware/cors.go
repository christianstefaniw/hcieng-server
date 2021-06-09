package middleware

import (
	"hciengserver/src/hciengserver"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	var origin string

	if hciengserver.DEBUG {
		origin = "http://localhost:3000"
	} else {
		origin = "https://hcieng.xyz"
	}

	return cors.New(cors.Config{
		AllowOrigins:     []string{origin},
		AllowMethods:     []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowHeaders:     []string{"X-Requested-With", "Content-Type"},
	})
}
