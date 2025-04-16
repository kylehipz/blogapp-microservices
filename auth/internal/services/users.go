package services

import (
	"context"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"golang.org/x/crypto/bcrypt"
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
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}

	createdUser, err := u.Queries.CreateUser(ctx, db.CreateUserParams{
		Username: username,
		Email:    email,
		Password: string(hashBytes),
	})
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}
