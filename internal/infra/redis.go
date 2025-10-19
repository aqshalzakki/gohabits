package infra

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
	ctx    *context.Context
}

func NewRedisClient(cfg *Config) *RedisClient {
	redisOpt := &redis.Options{
		Addr: cfg.Redis.Addr,
		DB:   cfg.Redis.DB,
	}

	if cfg.Redis.Password != "" {
		redisOpt.Password = cfg.Redis.Password
	}

	rdb := redis.NewClient(redisOpt)

	ctx := context.Background()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("❌ Failed to connect to Redis: %v", err)
	} else {
		log.Println("✅ Redis connected successfully")
	}

	return &RedisClient{
		client: rdb,
		ctx:    &ctx,
	}
}
