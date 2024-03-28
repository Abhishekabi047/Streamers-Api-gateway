package handlers

import (
	"api/pkg/client/interfaces"
	"api/pkg/models"
	"api/pkg/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
	Client interfaces.VideoClient
}

func NewVideoHandler(client interfaces.VideoClient) VideoHandler {
	return VideoHandler{
		Client: client,
	}
}

func (v *VideoHandler) UploadVideo(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists || userID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId not found in the context"})
		return
	}
	userid := userID.(int)
	fmt.Println("id", userid)
	title := c.PostForm("title")
	description := c.PostForm("description")
	category := c.PostForm("category")

	body := models.UploadVideo{
		Title:       title,
		Description: description,
		Category:    category,
		UserID:      userid,
	}

	file, err := c.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to find file",
			"error":   "error",
		})
		return
	}
	res, err1 := v.Client.UploadVideo(c.Request.Context(), file, body)
	if err1 != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "failed to find file",
			"error":   "error",
		})
		return
	}
	c.JSON(http.StatusOK, res)

}

// func (v *VideoHandler) StreamVideo(c *gin.Context) {
// 	filename:=c.Param("video_id")
// 	playlist:=c.Param("playlist")
// 	stream,err:=v.Client.StreamVideo(c.Request.Context(),filename,playlist)
// 	if err != nil{
// 		c.JSON(http.StatusInternalServerError,gin.H{
// 			"message":"failed to stream",
// 			"error":err.Error(),
// 		})
// 		return
// 	}
// 	for{
// 		resp,err:=stream.Recv()
// 		if err == io.EOF{
// 			break
// 		}
// 		if err != nil{
// 			c.JSON(http.StatusBadGateway,gin.H{
// 				"message":"error while recieving chunks"
// 				"error":err.Error(),
// 			})
// 		}
// 	}
// }

func (h *VideoHandler) FindUserVideo(ctx *gin.Context) {
	req := models.GetUserVideoRequest{}
	user := ctx.Query("userid")
	userId, _ := strconv.Atoi(user)
	req.User = userId

	if req.User == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "userid required",
		})
		return
	}
	res, err := h.Client.GetUserVideos(ctx, req)
	if err != nil {
		errmsg := utils.ExtractError(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": errmsg,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)

}

func (h *VideoHandler) FindAllVideo(ctx *gin.Context) {
	res, err := h.Client.FindAllVideo(ctx)
	if err != nil {
		errMsg := utils.ExtractError(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": errMsg,
			"error":   err,
		})
	}
	ctx.JSON(http.StatusOK, &res)
}

func (h *VideoHandler) GetVideoById(ctx *gin.Context) {
	videoId := ctx.Query("videoId")
	if videoId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "VideoId required",
		})
		return
	}
	data := models.GetVideoId{

		VideoId: videoId,
	}
	res, err := h.Client.GetVideoById(ctx, data)
	if err != nil {
		errMsg := utils.ExtractError(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)
}

func (h *VideoHandler) ArchiveVideo(ctx *gin.Context) {
	video := ctx.PostForm("videoid")
	body := models.ArchivedVideos{
		VideoId: video,
	}
	fmt.Println("body", body.VideoId)

	res, err := h.Client.ArchiveVideos(ctx, body)
	if err != nil {
		errMsg := utils.ExtractError(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)
}

func (h *VideoHandler) FindArchivedVideos(ctx *gin.Context) {
	req := models.GetUserVideoRequest{}
	user := ctx.Query("userId")

	userId, _ := strconv.Atoi(user)
	req.User = userId

	if req.User == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user name required",
		})
		return
	}
	res, err := h.Client.FindArchivedVideos(ctx, req)
	if err != nil {
		errMsg := utils.ExtractError(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)
}

func (v *VideoHandler) UploadClip(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists || userID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId not found in the context"})
		return
	}
	userid := userID.(int)
	fmt.Println("id", userid)
	title := c.PostForm("title")
	category := c.PostForm("category")

	body := models.UploadClip{
		Title:       title,
		Category:    category,
		UserID:      userid,
	}

	file, err := c.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to find file",
			"error":   "error",
		})
		return
	}
	res, err1 := v.Client.UploadClip(c.Request.Context(), file, body)
	if err1 != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "failed to find file",
			"error":   "error",
		})
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h *VideoHandler) FindUserClip(ctx *gin.Context) {
	req := models.GetUserVideoRequest{}
	user := ctx.Query("userid")
	userId, _ := strconv.Atoi(user)
	req.User = userId

	if req.User == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "userid required",
		})
		return
	}
	res, err := h.Client.GetUserClips(ctx, req)
	if err != nil {
		errmsg := utils.ExtractError(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": errmsg,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)

}


func (h *VideoHandler) FindAllClip(ctx *gin.Context) {
	res, err := h.Client.FindAllClip(ctx)
	if err != nil {
		errMsg := utils.ExtractError(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": errMsg,
			"error":   err,
		})
	}
	ctx.JSON(http.StatusOK, &res)
}


func (h *VideoHandler) GetClipById(ctx *gin.Context) {
	videoId := ctx.Query("videoId")
	if videoId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "VideoId required",
		})
		return
	}
	data := models.GetVideoId{

		VideoId: videoId,
	}
	res, err := h.Client.GetClipById(ctx, data)
	if err != nil {
		errMsg := utils.ExtractError(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)
}


func (h *VideoHandler) ArchiveClip(ctx *gin.Context) {
	video := ctx.PostForm("videoid")
	body := models.ArchivedVideos{
		VideoId: video,
	}
	fmt.Println("body", body.VideoId)

	res, err := h.Client.ArchiveClip(ctx, body)
	if err != nil {
		errMsg := utils.ExtractError(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)
}

func (h *VideoHandler) FindArchivedClips(ctx *gin.Context) {
	req := models.GetUserVideoRequest{}
	user := ctx.Query("userId")

	userId, _ := strconv.Atoi(user)
	req.User = userId

	if req.User == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user name required",
		})
		return
	}
	res, err := h.Client.FindArchivedClips(ctx, req)
	if err != nil {
		errMsg := utils.ExtractError(err.Error())
		log.Println(err)
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": errMsg,
			"error":   err,
		})
		return
	}
	ctx.JSON(http.StatusOK, &res)
}