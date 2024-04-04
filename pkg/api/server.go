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

func NewserverHttp(c *config.Config, Authhandler handlers.AuthHandler,videohandler handlers.VideoHandler,chatHandler handlers.ChatHandler) (*Server, error) {
	engine := gin.New()
	engine.Use(gin.Logger())
	routers.AuthRoutes(engine.Group("/"), Authhandler)
	routers.VideoRoutes(engine.Group("/video"), videohandler)
	routers.ChatRoutes(engine.Group("/"),chatHandler)
	
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
