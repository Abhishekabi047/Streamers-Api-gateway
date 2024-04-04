package routers

import (
	"api/pkg/api/handlers"
	"api/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func ChatRoutes(r *gin.RouterGroup, chatHandler handlers.ChatHandler) {
	r.Use(middlewares.CorsMiddleware)
	r.GET("/user/chathistory", chatHandler.ChatHistory)
	r.GET("/user/chatlist", chatHandler.ContactList)
	r.POST("/user/verify", chatHandler.VerifyContact)
}
