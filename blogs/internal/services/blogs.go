package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/types"
)

type BlogsService struct {
	Queries  *db.Queries
	dbClient db.DatabaseClient
}

func NewBlogsService(dbClient db.DatabaseClient) *BlogsService {
	return &BlogsService{dbClient: dbClient}
}

func (b *BlogsService) CreateBlog(
	ctx context.Context,
	authorId string,
	title string,
	content string,
) (*types.Blog, error) {
	createdBlog, err := b.dbClient.CreateBlog(ctx, authorId, title, content)
	if err != nil {
		return nil, err
	}

	return createdBlog, nil
}

func (b *BlogsService) GetBlog(ctx context.Context, blogId string) (*types.Blog, error) {
	foundBlog, err := b.dbClient.FindBlog(ctx, blogId)
	if err != nil {
		return nil, err
	}

	return foundBlog, nil
}

func (b *BlogsService) UpdateBlog(
	ctx context.Context,
	blog string,
	content string,
) (*db.Blog, error) {
	blogID, err := uuid.Parse(blog)
	if err != nil {
		return nil, err
	}

	createdBlog, err := b.Queries.UpdateBlog(ctx, db.UpdateBlogParams{
		ID:      blogID,
		Content: content,
	})
	if err != nil {
		return nil, err
	}

	return &createdBlog, err
}

func (b *BlogsService) DeleteBlog(ctx context.Context, blog string) error {
	blogID, err := uuid.Parse(blog)
	if err != nil {
		return err
	}

	err = b.Queries.DeleteBlog(ctx, blogID)
	if err != nil {
		return err
	}

	return nil
}
