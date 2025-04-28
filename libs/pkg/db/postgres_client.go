package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/kylehipz/blogapp-microservices/libs/internal/sqlcgen"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/types"
)

type PostgresClient struct {
	Queries *sqlcgen.Queries
}

func NewPostgresClient(conn *pgx.Conn) *PostgresClient {
	queries := sqlcgen.New(conn)

	return &PostgresClient{Queries: queries}
}

func (p *PostgresClient) GetHomeFeed(
	ctx context.Context,
	userId string,
	createdAt string,
	limit int32,
) ([]*types.Blog, error) {
	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}
	var t time.Time
	if createdAt == "now" {
		t = time.Now()
	} else {
		t, err = time.Parse(time.RFC3339, createdAt)
		if err != nil {
			return nil, err
		}
	}

	fetchedBlogs, err := p.Queries.GetHomeFeed(ctx, sqlcgen.GetHomeFeedParams{
		Follower:  parsedUserId,
		CreatedAt: t,
		Limit:     limit,
	})

	blogs := []*types.Blog{}

	for _, blog := range fetchedBlogs {
		blogs = append(blogs, &types.Blog{
			ID: blog.ID.String(),
			Author: &types.User{
				ID: userId,
			},
			Title:     "",
			Content:   blog.Content,
			CreatedAt: blog.CreatedAt.String(),
		})
	}

	return blogs, nil
}

func (p *PostgresClient) CreateBlog(
	ctx context.Context,
	authorId string,
	title string,
	content string,
) (*types.Blog, error) {
	parsedAuthorId, err := uuid.Parse(authorId)
	if err != nil {
		return nil, err
	}

	resultBlog, err := p.Queries.CreateBlog(ctx, sqlcgen.CreateBlogParams{
		Author:  parsedAuthorId,
		Content: content,
	})
	if err != nil {
		return nil, err
	}

	createdBlog := types.Blog{
		ID: resultBlog.ID.String(),
		Author: &types.User{
			ID: authorId,
		},
		Title:   resultBlog.Title,
		Content: resultBlog.Content,
	}

	return &createdBlog, err
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
