package models

import (
	"errors"
	"github.com/google/uuid"
	"pr0clone/util"
	"time"
)

type User struct {
	Id        string    `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	Enabled   bool      `json:"enabled" db:"enabled"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type UserDto struct {
	Id        string    `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Enabled   bool      `json:"enabled" db:"enabled"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

func NewUserDto(id string, username string, enabled bool, createdAt time.Time) *UserDto {
	return &UserDto{Id: id, Username: username, Enabled: enabled, CreatedAt: createdAt}
}

func NewUser(username string, password string) (*User, error) {
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, errors.New("could not hash password")
	}

	return &User{
		Id:        uuid.New().String(),
		Username:  username,
		Password:  hashedPassword,
		Enabled:   true,
		CreatedAt: time.Now(),
	}, nil
}

func NewUserFromUserReq(req UserReq) (*User, error) {
	return NewUser(req.Username, req.Password)
}

type UserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
