package services

import (
	"context"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/constants"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/pubsub"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/types"
)

type BlogsService struct {
	dbClient     db.DatabaseClient
	pubsubClient pubsub.PubSubClient
}

func NewBlogsService(dbClient db.DatabaseClient, pubsubClient pubsub.PubSubClient) *BlogsService {
	return &BlogsService{dbClient: dbClient, pubsubClient: pubsubClient}
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

	err = b.pubsubClient.Publish(ctx, constants.BLOG_CREATED, createdBlog)
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

	err = b.pubsubClient.Publish(ctx, constants.BLOG_UPDATED, updatedBlog)
	if err != nil {
		return nil, err
	}

	return updatedBlog, nil
}

func (b *BlogsService) DeleteBlog(ctx context.Context, blogId string, userId string) error {
	err := b.dbClient.DeleteBlog(ctx, blogId)
	if err != nil {
		return err
	}

	err = b.pubsubClient.Publish(ctx, constants.BLOG_DELETED, &types.Blog{
		ID:     blogId,
		Author: &types.User{ID: userId},
	})
	if err != nil {
		return err
	}

	return nil
}
