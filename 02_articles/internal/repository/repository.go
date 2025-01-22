package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Repositories struct {
	Articles   StorageArticles
	Categories StorageCategory
	Users      StorageUser
	UserRedis  RedisStorageUser
}

func NewRepositories(db *pgxpool.Pool, redis *redis.Client, log *zap.Logger) *Repositories {
	return &Repositories{
		Articles:   NewArticlesRepo(db, log),
		Categories: NewCategoryRepo(db, log),
		Users:      NewUserRepo(db, redis, log),
		UserRedis:  NewRedisStoreUser(redis),
	}
}
