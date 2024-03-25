package routers

import (
	"api/pkg/api/handlers"
	"api/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func VideoRoutes(r *gin.RouterGroup, videoHandler handlers.VideoHandler) {
	r.Use(middlewares.UserRetriveCookie)

	r.POST("/upload", videoHandler.UploadVideo)
	r.GET("/user-videos", videoHandler.FindUserVideo)
	r.GET("/user-archived-videos", videoHandler.FindArchivedVideos)
	r.GET("/list-all", videoHandler.FindAllVideo)
	r.POST("/archive-video", videoHandler.ArchiveVideo)
	r.GET("/get-by-id", videoHandler.GetVideoById)
}
