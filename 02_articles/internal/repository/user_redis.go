package repository

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorageUser interface {
	SetSession(ctx context.Context, RefreshToken string, sess Session) error
	GetSession(ctx context.Context, RefreshToken string) (*Session, error)
	DeleteSession(ctx context.Context, RefreshToken string) error
}

type RedisRepoUser struct {
	client *redis.Client
}

func NewRedisStoreUser(client *redis.Client) *RedisRepoUser {
	return &RedisRepoUser{client: client}
}

func (r *RedisRepoUser) SetSession(ctx context.Context, RefreshToken string, userID int, expireAt time.Duration) error {
	if err := r.client.HSet(ctx, RefreshToken, "UserID", userID, "ExpiresAt", expireAt).Err(); err != nil {
		return err
	}

	timeExpireSession := time.Now().Local().Add(10 * time.Second)

	r.client.ExpireAt(ctx, RefreshToken, timeExpireSession)

	return nil
}

func (r *RedisRepoUser) GetSession(ctx context.Context, RefreshToken string) (*Session, error) {

	sess := Session{}

	sessByRToken, err := r.client.HGetAll(ctx, RefreshToken).Result()
	if err != nil {
		return &sess, err
	}

	userID, ok := sessByRToken["UserID"]
	if !ok {
		return nil, err
	}

	userIDint, err := strconv.Atoi(userID)

	if err != nil {
		return nil, err
	}

	sess.UserID = userIDint

	return &sess, nil
}

func (r *RedisRepoUser) DeleteSession(ctx context.Context, RefreshToken string) error {
	err := r.client.Del(ctx, RefreshToken).Err()
	if err != nil {
		return err
	}

	return nil

}
