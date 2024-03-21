package api

import (
	"api/pkg/api/handlers"
	"api/pkg/api/routers"
	"api/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
	Port   string
}

func NewserverHttp(c *config.Config, Authhandler handlers.AuthHandler) (*Server, error) {
	engine := gin.New()
	engine.Use(gin.Logger())
	routers.RegisterRoutes(engine.Group("/"), Authhandler)
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"StatusCode": 404,
			"msg":        "invalid url",
		})
	})
	return &Server{
		Engine: engine,
		Port: c.Port,
	},nil
}

func(c *Server) Start() {
	c.Engine.Run(c.Port)
}
