package services

import (
	"context"
	"fmt"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/cache"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/pubsub"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/types"
)

type HomeFeedService struct {
	cacheClient  cache.CacheClient
	dbClient     db.DatabaseClient
	pubsubClient pubsub.PubSubClient
}

func NewHomeFeedService(
	dbClient db.DatabaseClient,
	cacheClient cache.CacheClient,
	pubsubClient pubsub.PubSubClient,
) *HomeFeedService {
	return &HomeFeedService{
		dbClient:     dbClient,
		cacheClient:  cacheClient,
		pubsubClient: pubsubClient,
	}
}

func (h *HomeFeedService) GetHomeFeed(
	ctx context.Context,
	userId string,
	createdAt string,
	limit int32,
) ([]*types.Blog, error) {
	cacheKey := h.generateHomeFeedCacheKey(userId, createdAt, limit)
	cachedResult, err := h.cacheClient.Get(ctx, cacheKey)
	if err == nil {
		cachedBlogs := cachedResult.([]*types.Blog)

		return cachedBlogs, nil
	} else {
		dbBlogs, err := h.dbClient.GetHomeFeed(ctx, userId, createdAt, limit)
		if err != nil {
			return nil, err
		}

		err = h.cacheClient.Set(ctx, cacheKey, dbBlogs)
		if err != nil {
			return nil, err
		}

		return dbBlogs, nil
	}
}

func (h *HomeFeedService) ListenToEvents(events []string) <-chan *pubsub.Message {
	messages, err := h.pubsubClient.Subscribe(events)
	if err != nil {
		panic(err)
	}

	return messages
}

func (h *HomeFeedService) generateHomeFeedCacheKey(
	user string,
	createdAt string,
	limit int32,
) string {
	return fmt.Sprintf("%s:%s:%d", user, createdAt, limit)
}
