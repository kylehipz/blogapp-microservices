package services

import "github.com/kylehipz/blogapp-microservices/libs/pkg/db"

type UsersService struct {
	Queries *db.Queries
}

func (u *UsersService) CreateUser(username string, email string) error {
	return nil
}
