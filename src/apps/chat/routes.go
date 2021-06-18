package chat

import (
	"hciengserver/src/apps/chat/controllers"
	"hciengserver/src/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterChatRoutes(router *gin.Engine) {
	api := router.Group("/chat")
	api.Use(middleware.WithAuth())

	api.POST("join/:id", controllers.JoinRoom)
	api.GET("info/:id/:lower_msg_bound/:upper_msg_bound", controllers.RoomInfo)
	api.GET("allrooms", controllers.AllUserRooms)
	api.GET("room/:id", controllers.ServeRoom)
	api.POST("newroom", controllers.NewRoom)
}
