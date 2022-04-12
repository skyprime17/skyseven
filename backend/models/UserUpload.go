package models

import (
	"github.com/google/uuid"
	"time"
)

type UserUpload struct {
	Id          string    `json:"id" db:"id"`
	UserId      string    `json:"userId" db:"user_id"`
	FileId      string    `json:"fileId" db:"file_id"`
	ThumbnailId string    `json:"thumbnailId" db:"thumbnail_id"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	Up          int64     `json:"up" db:"up"`
	Down        int64     `json:"down" db:"down"`
}

type UserUploadReq struct {
	Offset int64 `form:"offset" binding:"gte=0"`
	Limit  int64 `form:"limit"  binding:"max=150"`
}

func NewUserUpload(userId string, fileId string, thumbnailId string) *UserUpload {
	return &UserUpload{
		Id:          uuid.New().String(),
		UserId:      userId,
		FileId:      fileId,
		ThumbnailId: thumbnailId,
		CreatedAt:   time.Now(),
		Up:          0,
		Down:        0,
	}
}
