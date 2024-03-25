package interfaces

import (
	"api/pkg/models"
	"api/pkg/pb/video"
	"context"
	"mime/multipart"
)

type VideoClient interface {
	UploadVideo(context.Context, *multipart.FileHeader, models.UploadVideo) (*video.UploadVideoResponse, error)
	// StreamVideo(context.Context, string, string) (video.VideoService_StreamVideoClient, error)
	// FindAllVideo(context.Context) (*video.FindAllResponse, error)
	GetUserVideos(context.Context, models.GetUserVideoRequest) (*video.FindUserVideoResponse, error)
	FindAllVideo(context.Context) (*video.FindAllVideoResponse, error)
	GetVideoById(ctx context.Context, req models.GetVideoId) (*video.GetVideoByIdResponse, error)
	FindArchivedVideos(ctx context.Context, req models.GetUserVideoRequest) (*video.FindArchivedVideoByUserIdResponse, error)
	ArchiveVideos(ctx context.Context, req models.ArchivedVideos) (*video.ArchiveVideoResponse, error)
}
