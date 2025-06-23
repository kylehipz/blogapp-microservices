package db

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/kylehipz/blogapp-microservices/libs/internal/sqlcgen"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/errs"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/types"
)

type PostgresClient struct {
	queries *sqlcgen.Queries
}

func NewPostgresClient(conn *pgx.Conn) *PostgresClient {
	queries := sqlcgen.New(conn)

	return &PostgresClient{queries: queries}
}

func (p *PostgresClient) GetHomeFeed(
	ctx context.Context,
	userId string,
	createdAt string,
	limit int32,
) ([]*types.Blog, error) {
	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.ValidationError, err)
	}
	var t time.Time
	if createdAt == "now" {
		t = time.Now()
	} else {
		t, err = time.Parse(time.RFC3339, createdAt)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", errs.ApplicationError, err)
		}
	}

	fetchedBlogs, err := p.queries.GetHomeFeed(ctx, sqlcgen.GetHomeFeedParams{
		Follower:  parsedUserId,
		CreatedAt: t,
		Limit:     limit,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.DatabaseError, err)
	}

	blogs := []*types.Blog{}

	for _, blog := range fetchedBlogs {
		blogs = append(blogs, &types.Blog{
			ID: blog.ID.String(),
			Author: &types.User{
				ID: userId,
			},
			Title:     blog.Title,
			Content:   blog.Content,
			CreatedAt: blog.CreatedAt.Format(time.RFC3339),
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
		return nil, fmt.Errorf("%w: %v", errs.ValidationError, err)
	}

	resultBlog, err := p.queries.CreateBlog(ctx, sqlcgen.CreateBlogParams{
		Author:  parsedAuthorId,
		Title:   title,
		Content: content,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.DatabaseError, err)
	}

	createdBlog := types.Blog{
		ID: resultBlog.ID.String(),
		Author: &types.User{
			ID: authorId,
		},
		Title:     resultBlog.Title,
		Content:   resultBlog.Content,
		CreatedAt: resultBlog.CreatedAt.Format(time.RFC3339),
	}

	return &createdBlog, nil
}

func (p *PostgresClient) CreateUser(
	ctx context.Context,
	username string,
	email string,
	password string,
) (*types.User, error) {
	resultUser, err := p.queries.CreateUser(ctx, sqlcgen.CreateUserParams{
		Username: username,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.DatabaseError, err)
	}

	createdUser := types.User{
		ID:        resultUser.ID.String(),
		Username:  resultUser.Username,
		Email:     resultUser.Email,
		CreatedAt: resultUser.CreatedAt.Format(time.RFC3339),
	}

	return &createdUser, nil
}

func (p *PostgresClient) UpdateBlog(
	ctx context.Context,
	blogId string,
	title string,
	content string,
) (*types.Blog, error) {
	parsedBlogId, err := uuid.Parse(blogId)
	if err != nil {
		return nil, err
	}

	blogResult, err := p.queries.UpdateBlog(ctx, sqlcgen.UpdateBlogParams{
		ID:      parsedBlogId,
		Title:   title,
		Content: content,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.DatabaseError, err)
	}

	updatedBlog := &types.Blog{
		ID: blogResult.ID.String(),
		Author: &types.User{
			ID: blogResult.Author.String(),
		},
		Title:     blogResult.Title,
		Content:   blogResult.Content,
		CreatedAt: blogResult.CreatedAt.Format(time.RFC3339),
	}

	return updatedBlog, nil
}

func (p *PostgresClient) DeleteBlog(ctx context.Context, blogId string) error {
	parsedBlogId, err := uuid.Parse(blogId)
	if err != nil {
		return fmt.Errorf("%w: %v", errs.ValidationError, err)
	}

	err = p.queries.DeleteBlog(ctx, parsedBlogId)
	if err != nil {
		return fmt.Errorf("%w: %v", errs.DatabaseError, err)
	}

	return nil
}

func (p *PostgresClient) FindBlog(ctx context.Context, blogId string) (*types.Blog, error) {
	parsedBlogId, err := uuid.Parse(blogId)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.ValidationError, err)
	}

	resultBlog, err := p.queries.FindBlog(ctx, parsedBlogId)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.DatabaseError, err)
	}

	foundBlog := &types.Blog{
		ID: resultBlog.ID.String(),
		Author: &types.User{
			ID: resultBlog.Author.String(),
		},
		Title:     resultBlog.Title,
		Content:   resultBlog.Content,
		CreatedAt: resultBlog.CreatedAt.Format(time.RFC3339),
	}

	return foundBlog, nil
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
	user, err := p.queries.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.DatabaseError, err)
	}

	foundUser := &types.User{
		ID:        user.ID.String(),
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}

	return foundUser, nil
}

func (p *PostgresClient) CreateFollow(
	ctx context.Context,
	followerId string,
	followeeId string,
) (*types.Follow, error) {
	parsedFollowerId, err := uuid.Parse(followerId)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.ValidationError, err)
	}

	parsedFolloweeId, err := uuid.Parse(followeeId)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.ValidationError, err)
	}

	followResult, err := p.queries.CreateFollow(ctx, sqlcgen.CreateFollowParams{
		Follower: parsedFollowerId,
		Followee: parsedFolloweeId,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.DatabaseError, err)
	}

	createdFollow := &types.Follow{
		Follower: followResult.Follower.String(),
		Followee: followResult.Followee.String(),
	}

	return createdFollow, nil
}

func (p *PostgresClient) DeleteFollow(
	ctx context.Context,
	followerId string,
	followeeId string,
) error {
	parsedFollowerId, err := uuid.Parse(followerId)
	if err != nil {
		return fmt.Errorf("%w: %v", errs.ValidationError, err)
	}

	parsedFolloweeId, err := uuid.Parse(followeeId)
	if err != nil {
		return fmt.Errorf("%w: %v", errs.ValidationError, err)
	}

	err = p.queries.DeleteFollow(ctx, sqlcgen.DeleteFollowParams{
		Follower: parsedFollowerId,
		Followee: parsedFolloweeId,
	})
	if err != nil {
		return fmt.Errorf("%w: %v", errs.DatabaseError, err)
	}

	return nil
}

func (p *PostgresClient) GetFollowers(ctx context.Context, userId string) ([]uuid.UUID, error) {
	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.ValidationError, err)
	}

	followers, err := p.queries.FindFollowers(ctx, parsedUserId)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.DatabaseError, err)
	}

	return followers, nil
}
