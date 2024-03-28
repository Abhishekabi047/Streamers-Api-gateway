package routers

import (
	"api/pkg/api/handlers"
	"api/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func VideoRoutes(r *gin.RouterGroup, videoHandler handlers.VideoHandler) {
	r.GET("/list-all", middlewares.CorsMiddleware,videoHandler.FindAllVideo)
	r.GET("/get-by-id", middlewares.CorsMiddleware, videoHandler.GetVideoById)

	r.Use(middlewares.UserRetriveCookie,middlewares.CorsMiddleware)

	r.POST("/upload", videoHandler.UploadVideo)
	r.GET("/user-videos", videoHandler.FindUserVideo)
	r.GET("/user-archived-videos", videoHandler.FindArchivedVideos)
	
	r.POST("/archive-video", videoHandler.ArchiveVideo)
	
}
