package interfaces

import (
	"api/pkg/models"
	"api/pkg/pb/video"
	"context"
	"mime/multipart"
)

type VideoClient interface {
	ArchiveClip( context.Context,  models.ArchivedVideos) (*video.ArchiveClipResponse, error)
	ArchiveVideos( context.Context,  models.ArchivedVideos) (*video.ArchiveVideoResponse, error)
	FindAllClip( context.Context) (*video.FindAllClipResponse, error)
	FindAllVideo( context.Context) (*video.FindAllVideoResponse, error)
	FindArchivedClips( context.Context,  models.GetUserVideoRequest) (*video.FindArchivedClipByUserIdResponse, error)
	FindArchivedVideos( context.Context,  models.GetUserVideoRequest) (*video.FindArchivedVideoByUserIdResponse, error)
	GetClipById( context.Context,  models.GetVideoId) (*video.GetClipByIdResponse, error)
	GetUserClips( context.Context,  models.GetUserVideoRequest) (*video.FindUserClipResponse, error)
	GetUserVideos( context.Context,  models.GetUserVideoRequest) (*video.FindUserVideoResponse, error)
	GetVideoById( context.Context,  models.GetVideoId) (*video.GetVideoByIdResponse, error)
	UploadClip( context.Context,  *multipart.FileHeader,  models.UploadClip) (*video.UploadClipResponse, error)
	UploadVideo( context.Context,  *multipart.FileHeader,  models.UploadVideo) (*video.UploadVideoResponse, error)
}
