package cache

import "context"

type CacheClient interface {
	Get(ctx context.Context, key string) ([]string, error)
	Set(ctx context.Context, key string, value any) error
	RPush(ctx context.Context, key string, value ...any) error
}
