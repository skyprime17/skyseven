package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"pr0clone/errs"
	"time"
)

type JWTProvider struct {
	secretKey string
}

func NewJWTProvider(secret string) *JWTProvider {
	return &JWTProvider{secretKey: secret}
}

func (maker *JWTProvider) GenerateToken(username string, expTime time.Time) (string, error) {
	claims := NewJWTClaim(username, expTime)
	atTokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := atTokenClaim.SignedString([]byte(maker.secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (maker *JWTProvider) VerifyToken(token string) (*JWTClaim, error) {
	claims := &JWTClaim{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(maker.secretKey), nil
	})
	if err == jwt.ErrSignatureInvalid {
		return nil, errs.ErrInvalidToken
	}

	if !tkn.Valid {
		return nil, errs.ErrInvalidToken
	}
	return claims, nil
}
