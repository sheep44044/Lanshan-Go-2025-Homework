package cache

import (
	"awesomeProject1/homework07/config"
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func New(cfg *config.Config) (*RedisCache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &RedisCache{client: rdb}, nil
}

func (c *RedisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

func (c *RedisCache) SetWithRandomTTL(ctx context.Context, key string, value interface{}, baseTTL time.Duration) error {
	// 在基础TTL上增加±10%的随机浮动
	jitter := time.Duration(rand.Int63n(int64(baseTTL/5)) - int64(baseTTL/10))
	actualTTL := baseTTL + jitter

	// 确保TTL不会变成负数
	if actualTTL < 0 {
		actualTTL = baseTTL
	}

	return c.client.Set(ctx, key, value, actualTTL).Err()
}

func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *RedisCache) Del(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}
