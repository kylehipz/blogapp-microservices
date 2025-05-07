package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
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
	_, err := r.rdb.LPush(ctx, key, values...).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisClient) Get(ctx context.Context, key string) ([]string, error) {
	val, err := r.rdb.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (r *RedisClient) Set(ctx context.Context, key string, value any) error {
	marshaled, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err := r.rdb.Set(ctx, key, marshaled, time.Hour).Err(); err != nil {
		return err
	}

	return nil
}

func (r *RedisClient) Delete(ctx context.Context, key string) error {
	if _, err := r.rdb.Del(ctx, key).Result(); err != nil {
		return err
	}

	return nil
}
