package db

import (
	"context"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/types"
)

type PostgresClient struct {
	Queries *Queries
}

func (p *PostgresClient) GetHomeFeed(
	ctx context.Context,
	user string,
	createdAt string,
	limit int32,
) ([]*types.Blog, error) {
	return nil, nil
}

func (p *PostgresClient) CreateBlog(
	ctx context.Context,
	author string,
	title string,
	content string,
) (*types.Blog, error) {
	return nil, nil
}

func (p *PostgresClient) CreateUser(
	ctx context.Context,
	username string,
	email string,
	password string,
) (*types.User, error) {
	return nil, nil
}

func (p *PostgresClient) UpdateBlog(
	ctx context.Context,
	title string,
	content string,
) (*types.Blog, error) {
	return nil, nil
}

func (p *PostgresClient) DeleteBlog(ctx context.Context, blogId string) error {
	return nil
}

func (p *PostgresClient) FindBlog(ctx context.Context, blogId string) (*types.Blog, error) {
	return nil, nil
}

func (p *PostgresClient) FindUser(ctx context.Context, userId string) (*types.User, error) {
	return nil, nil
}

func (p *PostgresClient) FindUserByEmail(ctx context.Context, email string) (*types.User, error) {
	return nil, nil
}

func (p *PostgresClient) FindUserByUsername(
	ctx context.Context,
	username string,
) (*types.User, error) {
	return nil, nil
}

func (p *PostgresClient) FollowUser(
	ctx context.Context,
	follower string,
	followee string,
) error {
	return nil
}

func (p *PostgresClient) UnfollowUser(
	ctx context.Context,
	follower string,
	followee string,
) error {
	return nil
}
