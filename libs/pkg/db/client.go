package db

import (
	"context"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/types"
)

type DatabaseClient interface {
	GetHomeFeed(
		ctx context.Context,
		user string,
		createdAt string,
		limit int32,
	) ([]*types.Blog, error)
	CreateBlog(
		ctx context.Context,
		author string,
		title string,
		content string,
	) (*types.Blog, error)
	CreateUser(
		ctx context.Context,
		username string,
		email string,
		password string,
	) (*types.User, error)
	UpdateBlog(
		ctx context.Context,
		blogId string,
		title string,
		content string,
	) (*types.Blog, error)
	DeleteBlog(ctx context.Context, blogId string) error
	FindBlog(ctx context.Context, blogId string) (*types.Blog, error)
	FindUser(ctx context.Context, userId string) (*types.User, error)
	FindUserByEmail(ctx context.Context, email string) (*types.User, error)
	FindUserByUsername(ctx context.Context, username string) (*types.User, error)
	CreateFollow(ctx context.Context, follower string, followee string) (*types.Follow, error)
	DeleteFollow(ctx context.Context, follower string, followee string) error
}
