package auth

import (
	"api/pkg/client/interfaces"
	"api/pkg/config"
	"api/pkg/models"
	"fmt"

	"api/pkg/pb/video"
	"context"
	"io"
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
	fmt.Println("user", req.UserID)
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
	// userId, ok := ctx.Value("userId").(int)
	// if !ok {
	// 	log.Println("userId not found in context")
	// 	return nil, errors.New("login again")
	// }
	res, err := c.Server.GetVideoById(ctx, &video.GetVideoByIdRequest{
		VideoID: req.VideoId,
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


func (c *VideoClient) UploadClip(ctx context.Context, file *multipart.FileHeader, req models.UploadClip) (*video.UploadClipResponse, error) {
	uploadFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer uploadFile.Close()
	stream, err := c.Server.UploadClip(ctx)
	if err != nil {
		return nil, errr.Wrap(err, "failde to start upload stream")
	}
	chunkSize := 4096
	buffer := make([]byte, chunkSize)
	fmt.Println("user", req.UserID)
	for {
		n, err := uploadFile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if err := stream.Send(&video.UploadClipRequest{
			Filename:    file.Filename,
			Data:        buffer[:n],
			Title:       req.Title,
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


func (c *VideoClient) GetUserClips(ctx context.Context, req models.GetUserVideoRequest) (*video.FindUserClipResponse, error) {
	res, err := c.Server.FindUserClip(ctx, &video.FindUserClipRequest{
		Userid: int32(req.User),
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *VideoClient) FindAllClip(ctx context.Context) (*video.FindAllClipResponse, error) {
	res, err := c.Server.FindAllClip(ctx, &video.FindAllClipRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *VideoClient) GetClipById(ctx context.Context, req models.GetVideoId) (*video.GetClipByIdResponse, error) {
	// userId, ok := ctx.Value("userId").(int)
	// if !ok {
	// 	log.Println("userId not found in context")
	// 	return nil, errors.New("login again")
	// }
	res, err := c.Server.GetClipById(ctx, &video.GetClipByIdRequest{
		ClipId: req.VideoId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *VideoClient) ArchiveClip(ctx context.Context, req models.ArchivedVideos) (*video.ArchiveClipResponse, error) {
	res, err := c.Server.ArchiveClip(ctx, &video.ArchiveClipRequest{
		ClipId: req.VideoId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *VideoClient) FindArchivedClips(ctx context.Context, req models.GetUserVideoRequest) (*video.FindArchivedClipByUserIdResponse, error) {
	res, err := c.Server.FindArchivedClipByUserId(ctx, &video.FindArchivedClipByUserIdRequest{
		Userid: int32(req.User),
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

