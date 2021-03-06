package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"pr0clone/errs"
	"pr0clone/models"
	"pr0clone/services"
	"time"
)

type UserController struct {
	services.IAuthService
}

func (receiver *UserController) LogIn(ctx *gin.Context) {
	userReq := models.UserReq{}
	err := ctx.BindJSON(&userReq)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tokenPair, err := receiver.IAuthService.LogIn(ctx, userReq)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.SetCookie(
		"access_token",
		tokenPair["access_token"],
		60*15,
		"",
		"",
		false,
		true,
	)
	ctx.SetCookie(
		"refresh_token",
		tokenPair["refresh_token"],
		24*7*3600,
		"",
		"",
		false,
		true,
	)
}

func (receiver *UserController) Logout(ctx *gin.Context) {
	deleteAuthCookies(ctx)
}

func (receiver *UserController) Register(ctx *gin.Context) {
	userReq := models.UserReq{}
	err := ctx.BindJSON(&userReq)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	tokenPair, err := receiver.IAuthService.Register(ctx, userReq)

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.JSONP(http.StatusCreated, tokenPair)
}

func (receiver *UserController) Me(ctx *gin.Context) {
	user, err := receiver.IAuthService.CurrentUser(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.JSONP(http.StatusOK, user)

}

func (receiver *UserController) IsLoggedIn(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func (receiver *UserController) RefreshToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		err := errors.New("no Cookie provided")
		deleteAuthCookies(ctx)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errs.ErrorResponse(err))
		return
	}

	tokenMap, err := receiver.IAuthService.RefreshToken(refreshToken)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.SetCookie(
		"access_token",
		tokenMap["access_token"],
		time.Now().Add(15*time.Minute).Second(),
		"",
		"",
		false,
		true,
	)
}

func deleteAuthCookies(ctx *gin.Context) {
	ctx.SetCookie(
		"access_token",
		"",
		-1,
		"",
		"",
		false,
		true,
	)
	ctx.SetCookie(
		"refresh_token",
		"",
		-1,
		"",
		"",
		false,
		true,
	)
}
