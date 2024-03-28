package models

type UploadVideo struct {
	UserID      int    `json:"userid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

type GetUserVideoRequest struct {
	User int `json:"user"`
}

type GetVideoId struct {
	UserId  int `json:"user"`
	VideoId string
}

type ArchivedVideos struct {
	VideoId string `json:"videoid"`
}

type UploadClip struct {
	UserID      int    `json:"userid"`
	Title       string `json:"title"`
	Category    string `json:"category"`
}