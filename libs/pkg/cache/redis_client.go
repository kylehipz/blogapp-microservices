package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/kylehipz/blogapp-microservices/libs/pkg/errs"
)

type RedisClient struct {
	rdb *redis.Client
}

func NewRedisClient(rdb *redis.Client) *RedisClient {
	return &RedisClient{rdb: rdb}
}

func (r *RedisClient) RPush(ctx context.Context, key string, values ...any) error {
	_, err := r.rdb.RPush(ctx, key, values...).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisClient) LPush(ctx context.Context, key string, values ...any) error {
	if _, err := r.rdb.LPush(ctx, key, values...).Result(); err != nil {
		return fmt.Errorf("%w: %v", errs.CacheError, err)
	}

	return nil
}

func (r *RedisClient) Get(ctx context.Context, key string) ([]string, error) {
	val, err := r.rdb.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.CacheError, err)
	}

	return val, nil
}

func (r *RedisClient) Set(ctx context.Context, key string, value any) error {
	marshaled, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("%w: %v", errs.ValidationError, err)
	}

	if err := r.rdb.Set(ctx, key, marshaled, time.Hour).Err(); err != nil {
		return fmt.Errorf("%w: %v", errs.CacheError, err)
	}

	return nil
}

func (r *RedisClient) SetExpiration(ctx context.Context, key string, duration time.Duration) error {
	if _, err := r.rdb.Expire(ctx, key, duration).Result(); err != nil {
		return fmt.Errorf("%w: %v", errs.CacheError, err)
	}
	return nil
}

func (r *RedisClient) Delete(ctx context.Context, key string) error {
	if _, err := r.rdb.Del(ctx, key).Result(); err != nil {
		return fmt.Errorf("%w: %v", errs.CacheError, err)
	}

	return nil
}
