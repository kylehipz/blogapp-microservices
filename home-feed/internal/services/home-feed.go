package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/cache"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/db"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/pubsub"
	"github.com/kylehipz/blogapp-microservices/libs/pkg/types"

	"github.com/kylehipz/blogapp-microservices/home-feed/internal"
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
	cacheKey := h.generateHomeFeedCacheKey(userId)
	cachedHomeFeed, err := h.cacheClient.Get(ctx, cacheKey)
	if err != nil {
		return nil, err
	}

	parsedHomeFeedFromCache := h.unmarshalBlogs(cachedHomeFeed)

	requestedHomeFeedFromCache := h.getRequestedHomeFeed(parsedHomeFeedFromCache, createdAt, limit)

	if len(requestedHomeFeedFromCache) == 0 {
		// requested home feed is not yet in cache
		homeFeedFromDatabase, err := h.dbClient.GetHomeFeed(ctx, userId, createdAt, limit)
		if err != nil {
			return nil, err
		}

		if len(parsedHomeFeedFromCache) < internal.CACHE_LIMIT {
			if err = h.pushToCache(ctx, cacheKey, homeFeedFromDatabase); err != nil {
				return nil, err
			}

			if err = h.cacheClient.SetExpiration(ctx, cacheKey, 1*time.Hour); err != nil {
				return nil, err
			}
		}

		return homeFeedFromDatabase, nil
	}

	return requestedHomeFeedFromCache, nil
}

func (h *HomeFeedService) AddToCacheOfFollowers(
	ctx context.Context,
	blogFromEvent *types.Blog,
) error {
	followers, err := h.dbClient.GetFollowers(ctx, blogFromEvent.Author.ID)
	if err != nil {
		return err
	}

	for _, follower := range followers {
		cacheKey := h.generateHomeFeedCacheKey(follower.String())
		cachedHomeFeed, err := h.cacheClient.Get(ctx, cacheKey)
		if err != nil {
			return err
		}

		if len(cachedHomeFeed) == 0 {
			// no cache for this user
			continue
		}

		bytes, err := json.Marshal(blogFromEvent)
		if err != nil {
			return err
		}
		if err = h.cacheClient.LPush(ctx, cacheKey, bytes); err != nil {
			return err
		}
	}

	return nil
}

func (h *HomeFeedService) UpdateInCacheOfFollowers(
	ctx context.Context,
	blogFromEvent *types.Blog,
) error {
	followers, err := h.dbClient.GetFollowers(ctx, blogFromEvent.Author.ID)
	if err != nil {
		return err
	}

	for _, follower := range followers {
		cacheKey := h.generateHomeFeedCacheKey(follower.String())
		cachedHomeFeed, err := h.cacheClient.Get(ctx, cacheKey)
		if err != nil {
			return err
		}

		if len(cachedHomeFeed) == 0 {
			// no cache for this user
			continue
		}

		updatedHomeFeed := []*types.Blog{}

		for _, blogStr := range cachedHomeFeed {
			blog := &types.Blog{}
			err := json.Unmarshal([]byte(blogStr), blog)
			if err != nil {
				continue
			}

			if blog.ID == blogFromEvent.ID {
				updatedHomeFeed = append(updatedHomeFeed, blogFromEvent)
			} else {
				updatedHomeFeed = append(updatedHomeFeed, blog)
			}
		}

		if err = h.cacheClient.Delete(ctx, cacheKey); err != nil {
			continue
		}

		if err = h.pushToCache(ctx, cacheKey, updatedHomeFeed); err != nil {
			continue
		}
	}

	return nil
}

func (h *HomeFeedService) DeleteInCacheOfFollowers(
	ctx context.Context,
	blogFromEvent *types.Blog,
) error {
	followers, err := h.dbClient.GetFollowers(ctx, blogFromEvent.Author.ID)
	if err != nil {
		return err
	}

	for _, follower := range followers {
		cacheKey := h.generateHomeFeedCacheKey(follower.String())
		cachedHomeFeed, err := h.cacheClient.Get(ctx, cacheKey)
		if err != nil {
			return err
		}

		if len(cachedHomeFeed) == 0 {
			// no cache for this user
			continue
		}

		updatedHomeFeed := []*types.Blog{}

		for _, blogStr := range cachedHomeFeed {
			blog := &types.Blog{}
			err := json.Unmarshal([]byte(blogStr), blog)
			if err != nil {
				continue
			}

			if blog.ID == blogFromEvent.ID {
				continue
			} else {
				updatedHomeFeed = append(updatedHomeFeed, blog)
			}
		}

		if err = h.cacheClient.Delete(ctx, cacheKey); err != nil {
			continue
		}

		if err = h.pushToCache(ctx, cacheKey, updatedHomeFeed); err != nil {
			continue
		}
	}

	return nil
}

func (h *HomeFeedService) ListenToEvents(events []string) <-chan *pubsub.Message {
	messages, err := h.pubsubClient.Subscribe(events)
	if err != nil {
		panic(err)
	}

	return messages
}

func (h *HomeFeedService) generateHomeFeedCacheKey(user string) string {
	return fmt.Sprintf("%s:home-feed", user)
}

func (h *HomeFeedService) unmarshalBlogs(homeFeedStr []string) []*types.Blog {
	blogs := []*types.Blog{}
	for _, blogStr := range homeFeedStr {
		blog := &types.Blog{}

		err := json.Unmarshal([]byte(blogStr), blog)
		if err != nil {
			log.Println("Error decoding json", blogStr)
			continue
		}

		blogs = append(blogs, blog)
	}

	return blogs
}

func (h *HomeFeedService) getRequestedHomeFeed(
	homeFeed []*types.Blog,
	createdAt string,
	limit int32,
) []*types.Blog {
	requestedHomeFeed := []*types.Blog{}
	var requestedCreatedAt time.Time
	if createdAt == "now" {
		requestedCreatedAt = time.Now()
	} else {
		t, err := time.Parse(time.RFC3339, createdAt)
		requestedCreatedAt = t
		if err != nil {
			return nil
		}
	}

	for index, blog := range homeFeed {
		if int32(index) == limit {
			break
		}
		blogCreatedAt, err := time.Parse(time.RFC3339, blog.CreatedAt)
		if err != nil {
			log.Println("Error parsing date")
			continue
		}

		if blogCreatedAt.Before(requestedCreatedAt) {
			requestedHomeFeed = append(requestedHomeFeed, blog)
		}
	}

	return requestedHomeFeed
}

func (h *HomeFeedService) pushToCache(
	ctx context.Context,
	cacheKey string,
	homeFeed []*types.Blog,
) error {
	var values []any

	for _, blog := range homeFeed {
		b, _ := json.Marshal(blog)
		values = append(values, b)
	}

	if len(values) == 0 {
		return nil
	}

	err := h.cacheClient.RPush(ctx, cacheKey, values...)

	return err
}
