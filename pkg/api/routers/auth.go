package routers

import (
	"api/pkg/api/handlers"
	"api/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup,authHandler handlers.AuthHandler){
	r.Use(middlewares.CorsMiddleware)

	r.POST("/user/signup",authHandler.Signup)
	r.POST("/user/otp",authHandler.Otp)
	r.POST("/user/login",middlewares.CorsMiddleware,authHandler.Login)
}