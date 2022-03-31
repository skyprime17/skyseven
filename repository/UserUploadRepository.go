package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"pr0clone/models"
)

type IUserUploadRepository interface {
	Create(ctx context.Context, user *models.UserUpload) error
	GetFilesByUserId(ctx context.Context, username string) ([]models.UserUpload, error)
	GetTopUploads(ctx context.Context, uploadReq models.UserUploadReq) ([]models.UserUpload, error)
	GetPostById(ctx context.Context, uploadId string) (models.UserUpload, error)
	GetNewestUploads(ctx context.Context, uploadReq models.UserUploadReq) ([]models.UserUpload, error)
}

type UserUploadRepository struct {
	client *sqlx.DB
}

func NewUserUploadRepository(client *sqlx.DB) *UserUploadRepository {
	return &UserUploadRepository{client: client}
}

func (receiver *UserUploadRepository) Create(ctx context.Context, userUpload *models.UserUpload) error {
	_, err := receiver.client.NamedExecContext(ctx,
		`INSERT INTO user_upload(id, user_id, thumbnail_id, file_id, created_at, up, down)
			   VALUES (:id, :user_id, :thumbnail_id, :file_id , :created_at, :up, :down)`,
		userUpload)
	if err != nil {
		return err
	}
	return nil
}

func (receiver *UserUploadRepository) GetFilesByUserId(ctx context.Context, username string) ([]models.UserUpload, error) {
	return nil, nil
}

func (receiver *UserUploadRepository) GetPostById(ctx context.Context, uploadId string) (models.UserUpload, error) {
	var up models.UserUpload
	err := receiver.client.GetContext(ctx, &up, "SELECT * FROM user_upload WHERE id=?", uploadId)
	return up, err
}

func (receiver *UserUploadRepository) GetTopUploads(ctx context.Context, uploadReq models.UserUploadReq) ([]models.UserUpload, error) {
	var ups []models.UserUpload
	err := receiver.client.SelectContext(ctx, &ups,
		"SELECT * FROM user_upload ORDER BY up DESC, down DESC LIMIT ? OFFSET ?",
		uploadReq.Limit, uploadReq.Offset)
	return ups, err
}

func (receiver *UserUploadRepository) GetNewestUploads(ctx context.Context, uploadReq models.UserUploadReq) ([]models.UserUpload, error) {
	var ups []models.UserUpload
	err := receiver.client.SelectContext(ctx, &ups,
		"SELECT * FROM user_upload ORDER BY created_at LIMIT ? OFFSET ?",
		uploadReq.Limit, uploadReq.Offset)
	return ups, err
}
