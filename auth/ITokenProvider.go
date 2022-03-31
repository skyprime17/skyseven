package auth

import "time"

type ITokenProvider interface {
	GenerateToken(username string, expTime time.Time) (string, error)
	VerifyToken(token string) (*JWTClaim, error)
}
