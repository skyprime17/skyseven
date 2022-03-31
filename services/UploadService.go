package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
	"path/filepath"
	"pr0clone/models"
	"pr0clone/repository"
	"strings"
)

type IUserUploadService interface {
	UploadFile(ctx *gin.Context, file *multipart.FileHeader) error
	GetTopPosts(ctx *gin.Context, uploadReq models.UserUploadReq) ([]models.UserUpload, error)
	GetNewestPosts(ctx *gin.Context, uploadReq models.UserUploadReq) ([]models.UserUpload, error)
	GetPostById(context *gin.Context, uploadId string) (models.UserUpload, error)
}

type UserUploadService struct {
	UserUploadRepository repository.IUserUploadRepository
	AuthService          IAuthService
}

func NewUserUploadService(userUploadRepository repository.IUserUploadRepository, authService IAuthService) *UserUploadService {
	return &UserUploadService{UserUploadRepository: userUploadRepository, AuthService: authService}
}

func (receiver UserUploadService) UploadFile(ctx *gin.Context, file *multipart.FileHeader) error {
	// Retrieve file information
	//TODO restrict viable file extensions so it cant get exploited
	//TODO thumbnail generation
	extension := strings.ToLower(filepath.Ext(file.Filename))
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String()

	currentUser, err := receiver.AuthService.CurrentUser(ctx)

	if err != nil {
		return err
	}

	fileName := newFileName + extension
	userUpload := models.NewUserUpload(currentUser.Id, fileName, "")

	// The file is received, so let's save it
	if err := ctx.SaveUploadedFile(file, ".\\test\\"+fileName); err != nil {
		return err
	}

	if err := receiver.UserUploadRepository.Create(ctx, userUpload); err != nil {
		fmt.Println("here")
		return err
	}

	return nil

}

func (receiver UserUploadService) GetTopPosts(ctx *gin.Context, uploadReq models.UserUploadReq) ([]models.UserUpload, error) {
	return receiver.UserUploadRepository.GetTopUploads(ctx, uploadReq)
}

func (receiver UserUploadService) GetNewestPosts(ctx *gin.Context, uploadReq models.UserUploadReq) ([]models.UserUpload, error) {
	return receiver.UserUploadRepository.GetNewestUploads(ctx, uploadReq)
}

func (receiver UserUploadService) GetPostById(ctx *gin.Context, uploadId string) (models.UserUpload, error) {
	return receiver.UserUploadRepository.GetPostById(ctx, uploadId)
}
