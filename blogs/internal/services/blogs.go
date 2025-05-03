package services

import (
	"context"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/pubsub"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/types"
)

type BlogsService struct {
	dbClient       db.DatabaseClient
	rabbitMQClient *pubsub.RabbitMQClient
}

func NewBlogsService(
	dbClient db.DatabaseClient,
	rabbitMQClient *pubsub.RabbitMQClient,
) *BlogsService {
	return &BlogsService{
		dbClient:       dbClient,
		rabbitMQClient: rabbitMQClient,
	}
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

	err = b.rabbitMQClient.Publish(ctx, "blog.created", createdBlog)
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
	blogId string,
	title string,
	content string,
) (*types.Blog, error) {
	updatedBlog, err := b.dbClient.UpdateBlog(ctx, blogId, title, content)
	if err != nil {
		return nil, err
	}

	err = b.rabbitMQClient.Publish(ctx, "blog.updated", updatedBlog)
	if err != nil {
		return nil, err
	}

	return updatedBlog, nil
}

func (b *BlogsService) DeleteBlog(ctx context.Context, blogId string) error {
	err := b.dbClient.DeleteBlog(ctx, blogId)
	if err != nil {
		return err
	}

	err = b.rabbitMQClient.Publish(ctx, "blog.updated", map[string]bool{
		"success": true,
	})
	if err != nil {
		return err
	}

	return nil
}
