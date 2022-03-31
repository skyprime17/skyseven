package errs

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

var (
	MalformedAuthHeaderErr = errors.New("authorization header malformed")
	ErrInvalidToken        = errors.New("token is invalid")
	ErrExpiredToken        = errors.New("token has expired")
	InvalidTokenError      = errors.New("invalid Token")
)
