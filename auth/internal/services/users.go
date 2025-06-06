package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/types"
	"golang.org/x/crypto/bcrypt"
)

var InvalidCredentialsError = errors.New("Invalid Credentials")

type UsersService struct {
	dbClient db.DatabaseClient
}

func NewUsersService(dbClient db.DatabaseClient) *UsersService {
	return &UsersService{dbClient: dbClient}
}

func (u *UsersService) CreateUser(
	ctx context.Context,
	username string,
	email string,
	password string,
) (*types.User, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}

	createdUser, err := u.dbClient.CreateUser(ctx, username, email, string(hashBytes))
	if err != nil {
		fmt.Println(strings.Contains(err.Error(), "duplicate"))
		return nil, err
	}

	return createdUser, nil
}

func (u *UsersService) Login(
	ctx context.Context,
	username string,
	password string,
) (string, error) {
	// check if user exists
	user, err := u.dbClient.FindUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	// verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", InvalidCredentialsError
	}

	// generate jwt
	claims := &types.JwtCustomClaims{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
