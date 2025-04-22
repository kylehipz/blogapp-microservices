package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"github.com/redis/go-redis/v9"
)

type HomeFeedService struct {
	Queries     *db.Queries
	RedisClient *redis.Client
}

func (h *HomeFeedService) GetHomeFeed(
	ctx context.Context,
	user string,
	createdAt string,
	page int32,
	limit int32,
) ([]db.Blog, error) {
	fmt.Println("user", user)
	userID, err := uuid.Parse(user)
	if err != nil {
		return nil, err
	}

	// convert timestamp to time.Time
	t, err := time.Parse(time.RFC3339, createdAt)
	if err != nil {
		return nil, err
	}

	// fetch from cache
	cacheKey := h.generateHomeFeedCacheKey(user, page, limit)
	val, err := h.RedisClient.Get(ctx, cacheKey).Result()
	if err != nil {
		// fetch from database
		blogs, err := h.Queries.GetHomeFeed(ctx, db.GetHomeFeedParams{
			Follower:  userID,
			CreatedAt: t,
			Limit:     limit,
		})
		if err != nil {
			return nil, err
		}

		blogsMarshaled, err := json.Marshal(&blogs)
		if err != nil {
			return nil, err
		}

		err = h.RedisClient.Set(ctx, cacheKey, blogsMarshaled, 0).Err()
		if err != nil {
			return nil, err
		}

		return blogs, nil
	}

	var cachedHomeFeed []db.Blog

	err = json.Unmarshal([]byte(val), &cachedHomeFeed)
	if err != nil {
		return nil, err
	}

	return cachedHomeFeed, nil
}

func (h *HomeFeedService) generateHomeFeedCacheKey(user string, page int32, limit int32) string {
	return fmt.Sprintf("%s:%d:%d", user, page, limit)
}
