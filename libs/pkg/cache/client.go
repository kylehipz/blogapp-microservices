package cache

import (
	"context"
	"time"
)

type CacheClient interface {
	Get(ctx context.Context, key string) ([]string, error)
	Set(ctx context.Context, key string, value any) error
	SetExpiration(ctx context.Context, key string, duration time.Duration) error
	RPush(ctx context.Context, key string, value ...any) error
	LPush(ctx context.Context, key string, value ...any) error
	Delete(ctx context.Context, key string) error
}
