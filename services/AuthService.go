package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	auth2 "pr0clone/auth"
	"pr0clone/models"
	"pr0clone/repository"
	"pr0clone/util"
	"time"
)

type IAuthService interface {
	LogIn(ctx *gin.Context, userReq models.UserReq) (map[string]string, error)
	Register(ctx *gin.Context, userReq models.UserReq) (map[string]string, error)
	RefreshToken(refreshToken string) (map[string]string, error)
	CurrentUser(ctx *gin.Context) (*models.UserDto, error)
}

type AuthService struct {
	UserRepository repository.IUserRepository
	TokenMaker     auth2.ITokenProvider
}

func NewAuthService(userRepository repository.IUserRepository, tokenMaker auth2.ITokenProvider) *AuthService {
	return &AuthService{UserRepository: userRepository, TokenMaker: tokenMaker}
}

func (receiver *AuthService) LogIn(ctx *gin.Context, userReq models.UserReq) (map[string]string, error) {
	user, err := receiver.UserRepository.GetUserByUsername(ctx, userReq.Username)

	if err != nil {
		return nil, errors.New("unknown user")
	}

	if !util.CheckPasswordHash(userReq.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	accessToken, err := receiver.TokenMaker.GenerateToken(user.Username, time.Now().Add(15*time.Minute))
	if err != nil {
		return nil, errors.New("couldnt generate access token for user")
	}

	refreshToken, err := receiver.TokenMaker.GenerateToken(user.Username, time.Now().Add(24*7*time.Hour))
	if err != nil {
		return nil, errors.New("couldnt generate refresh token for user")
	}
	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

func (receiver *AuthService) Register(ctx *gin.Context, userReq models.UserReq) (map[string]string, error) {

	user, err := models.NewUserFromUserReq(userReq)

	if err != nil {
		return nil, errors.New("could not create new user obj")
	}

	err = receiver.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.New("user already existing")
	}

	accessToken, err := receiver.TokenMaker.GenerateToken(user.Username, time.Now().Add(1*time.Minute))
	if err != nil {
		return nil, errors.New("couldnt generate access token for user")
	}

	refreshToken, err := receiver.TokenMaker.GenerateToken(user.Username, time.Now().Add(24*30*time.Hour))
	if err != nil {
		return nil, errors.New("couldnt generate refresh token for user")
	}
	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

func (receiver *AuthService) RefreshToken(refreshToken string) (map[string]string, error) {
	claims, err := receiver.TokenMaker.VerifyToken(refreshToken)
	if err != nil {
		return nil, err
	}

	accessToken, err := receiver.TokenMaker.GenerateToken(claims.Username, time.Now().Add(1*time.Minute))
	if err != nil {
		return nil, errors.New("couldnt generate access token for user")
	}
	return map[string]string{
		"access_token": accessToken,
	}, nil
}

func (receiver *AuthService) CurrentUser(ctx *gin.Context) (*models.UserDto, error) {
	userClaims, _ := ctx.Get("user")
	username := userClaims.(*auth2.JWTClaim).Username
	user, err := receiver.UserRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	userDto := models.NewUserDto(user.Id, user.Username, user.Enabled, user.CreatedAt)
	return userDto, nil
}
