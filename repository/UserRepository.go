package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"pr0clone/models"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, userId string) (models.User, error)
}

type UserRepository struct {
	client *sqlx.DB
}

func NewUserRepository(client *sqlx.DB) *UserRepository {
	return &UserRepository{client: client}
}

func (receiver *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	_, err := receiver.client.NamedExecContext(ctx,
		`INSERT INTO user(id, username, password, enabled, created_at) VALUES (:id, :username, :password , :enabled, :created_at)`,
		user)
	if err != nil {
		return err
	}
	return nil
}

func (receiver *UserRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	user := models.User{}
	err := receiver.client.GetContext(ctx, &user, "SELECT * FROM user WHERE username = ?", username)
	fmt.Println(user, err)
	return user, err
}
