package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"pr0clone/auth"
	"pr0clone/errs"
)

var (
	authHeaderKey = "Authorization"
	authType      = "jwt"
	authClaimsKey = "user"
)

func AuthMiddleware(tokenMaker auth.ITokenProvider) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/*
			authHeader := ctx.GetHeader(authHeaderKey)

			if len(authHeader) == 0 {
				err := errors.New("authorization header is not provided")
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, errs.ErrorResponse(err))
				return
			}

			authSplit := strings.Split(authHeader, authType)
			if len(authSplit) != 2 {
				err := errors.New("authorization header malformed")
				ctx.AbortWithStatusJSON(http.StatusBadRequest, errs.ErrorResponse(err))
			}

			authStr := strings.TrimSpace(authSplit[1])
			claims, err := tokenMaker.VerifyToken(authStr)

			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, errs.ErrorResponse(err))
			}

			ctx.Set(authClaimsKey, claims)
			ctx.Next()
		*/
		accessToken, err := ctx.Cookie("access_token")
		if err != nil {
			err := errors.New("no Cookie provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errs.ErrorResponse(err))
			return
		}
		claims, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errs.ErrorResponse(err))
			return
		}
		ctx.Set(authClaimsKey, claims)
		ctx.Next()

	}
}
