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
