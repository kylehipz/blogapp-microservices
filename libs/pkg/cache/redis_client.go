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

func (r *RedisClient) Get(ctx context.Context, key string) (any, error) {
	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var parsed any

	err = json.Unmarshal([]byte(val), parsed)
	if err != nil {
		return nil, err
	}

	return parsed, nil
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
