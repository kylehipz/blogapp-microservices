package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
)

type BlogsService struct {
	Queries *db.Queries
}

func (b *BlogsService) CreateBlog(
	ctx context.Context,
	author string,
	content string,
) (*db.Blog, error) {
	authorID, err := uuid.Parse(author)
	if err != nil {
		return nil, err
	}

	createdBlog, err := b.Queries.CreateBlog(ctx, db.CreateBlogParams{
		Author:  authorID,
		Content: content,
	})
	if err != nil {
		return nil, err
	}

	return &createdBlog, err
}

func (b *BlogsService) GetBlog(ctx context.Context, blog string) (*db.Blog, error) {
	blogID, err := uuid.Parse(blog)
	if err != nil {
		return nil, err
	}

	existingBlog, err := b.Queries.FindBlog(ctx, blogID)
	if err != nil {
		return nil, err
	}

	return &existingBlog, nil
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
