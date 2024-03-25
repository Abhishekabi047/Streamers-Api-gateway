package auth

import (
	"api/pkg/client/interfaces"
	"api/pkg/config"
	"api/pkg/models"
	"fmt"

	"api/pkg/pb/video"
	"context"
	"errors"
	"io"
	"log"
	"mime/multipart"

	errr "github.com/pkg/errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type VideoClient struct {
	Server video.VideoServiceClient
}

func InitVideoClient(c *config.Config) (video.VideoServiceClient, error) {
	cc, err := grpc.Dial(c.VideoService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return video.NewVideoServiceClient(cc), nil
}

func NewVideoClient(server video.VideoServiceClient) interfaces.VideoClient {
	return &VideoClient{
		Server: server,
	}
}

func (c *VideoClient) UploadVideo(ctx context.Context, file *multipart.FileHeader, req models.UploadVideo) (*video.UploadVideoResponse, error) {
	uploadFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer uploadFile.Close()
	stream, err := c.Server.UploadVideo(ctx)
	if err != nil {
		return nil, errr.Wrap(err, "failde to start upload stream")
	}
	chunkSize := 4096
	buffer := make([]byte, chunkSize)
	fmt.Println("user",req.UserID)
	for {
		n, err := uploadFile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if err := stream.Send(&video.UploadVideoRequest{
			Filename:    file.Filename,
			Data:        buffer[:n],
			Title:       req.Title,
			Description: req.Description,
			UserId:      int32(req.UserID),
			Category:    req.Category,
		}); err != nil {
			return nil, err
		}

	}
	resposne, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}
	return resposne, nil
}

// func (c *VideoClient) StreamVideo(ctx context.Context, filename, playlist string) (video.VideoService_StreamVideoClient, error) {
// 	res, err := c.Server.StreamVideo(ctx, &video.StreamVideoRequest{
// 		Videoid:  filename,
// 		Playlist: playlist,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

// func (c *VideoClient) FindAllVideo(ctx context.Context) (*video.FindAllResponse, error) {
// 	res, err := c.Server.FindAllVideo(ctx, &video.FindAllRequest{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

func (c *VideoClient) GetUserVideos(ctx context.Context, req models.GetUserVideoRequest) (*video.FindUserVideoResponse, error) {
	res, err := c.Server.FindUserVideo(ctx, &video.FindUserVideoRequest{
		Userid: int32(req.User),
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *VideoClient) FindAllVideo(ctx context.Context) (*video.FindAllVideoResponse, error) {
	res, err := c.Server.FindAllVideo(ctx, &video.FindAllVideoRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *VideoClient) GetVideoById(ctx context.Context, req models.GetVideoId) (*video.GetVideoByIdResponse, error) {
	userId, ok := ctx.Value("userId").(int)
	if !ok {
		log.Println("userId not found in context")
		return nil, errors.New("login again")
	}
	res, err := c.Server.GetVideoById(ctx, &video.GetVideoByIdRequest{
		VideoID: req.VideoId,
		UserId:  int32(userId),
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *VideoClient) ArchiveVideos(ctx context.Context, req models.ArchivedVideos) (*video.ArchiveVideoResponse, error) {
	res, err := c.Server.ArchiveVideo(ctx, &video.ArchiveVideoRequest{
		VideoId: req.VideoId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *VideoClient) FindArchivedVideos(ctx context.Context, req models.GetUserVideoRequest) (*video.FindArchivedVideoByUserIdResponse, error) {
	res, err := c.Server.FindArchivedVideoByUserId(ctx, &video.FindArchivedVideoByUserIdRequest{
		Userid: int32(req.User),
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
