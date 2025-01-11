package redis_storage

import (
	"context"

	"github.com/gMerl1on/parsers_articles/02_articles/configs"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg configs.ConfigRedis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.AddrRedis,
		Password: cfg.PasswordRedis,
		DB:       cfg.DBRedis,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// redis:
//   addr_redis: "redis:6379"
//   db_redis: "0"
