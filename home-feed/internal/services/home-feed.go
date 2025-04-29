package services

import (
	"context"
	"fmt"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/types"
	"github.com/redis/go-redis/v9"
)

type HomeFeedService struct {
	redisClient *redis.Client
	dbClient    db.DatabaseClient
}

func (h *HomeFeedService) GetHomeFeed(
	ctx context.Context,
	userId string,
	createdAt string,
	limit int32,
) ([]*types.Blog, error) {
	blogs, err := h.dbClient.GetHomeFeed(ctx, userId, createdAt, limit)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (h *HomeFeedService) generateHomeFeedCacheKey(
	user string,
	createdAt string,
	limit int32,
) string {
	return fmt.Sprintf("%s:%s:%d", user, createdAt, limit)
}
