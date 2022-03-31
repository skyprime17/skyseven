package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pr0clone/models"
	"pr0clone/services"
)

type UserUploadController struct {
	services.IUserUploadService
}

func (receiver *UserUploadController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	if err := receiver.IUserUploadService.UploadFile(c, file); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
}

func (receiver *UserUploadController) RemoveFileById(ctx *gin.Context) {

}

// GetTopPosts returns the top posts, max 2, with offset
func (receiver *UserUploadController) GetTopPosts(ctx *gin.Context) {
	var uploadReq models.UserUploadReq
	if err := ctx.ShouldBindQuery(&uploadReq); err != nil {
		ctx.AbortWithStatus(400)
		return
	}
	posts, err := receiver.IUserUploadService.GetTopPosts(ctx, uploadReq)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}
	ctx.JSONP(http.StatusOK, posts)
}

func (receiver *UserUploadController) GetNewestPosts(ctx *gin.Context) {
	var uploadReq models.UserUploadReq
	if err := ctx.ShouldBindQuery(&uploadReq); err != nil {
		ctx.AbortWithStatus(400)
		return
	}
	posts, err := receiver.IUserUploadService.GetNewestPosts(ctx, uploadReq)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}
	ctx.JSONP(http.StatusOK, posts)
}

func (receiver *UserUploadController) GetPostById(ctx *gin.Context) {
	uploadId := ctx.Param("itemId")
	if uploadId == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	post, err := receiver.IUserUploadService.GetPostById(ctx, uploadId)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}
	ctx.JSONP(http.StatusOK, post)
}
