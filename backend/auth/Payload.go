package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

type JWTClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewJWTClaim(username string, expTime time.Time) *JWTClaim {
	return &JWTClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
			Id:        uuid.New().String(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "skyseven",
		}}
}
