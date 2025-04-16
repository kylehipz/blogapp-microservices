package services

import (
	"context"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
)

type UsersService struct {
	Queries *db.Queries
}

func (u *UsersService) CreateUser(
	ctx context.Context,
	username string,
	email string,
	password string,
) (*db.User, error) {
	createdUser, err := u.Queries.CreateUser(ctx, db.CreateUserParams{
		Username: username,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}
