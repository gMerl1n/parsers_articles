package service

import (
	"context"
	"fmt"
	"time"

	"github.com/gMerl1on/parsers_articles/02_articles/internal/domain"
	"github.com/gMerl1on/parsers_articles/02_articles/internal/repository"
	er "github.com/gMerl1on/parsers_articles/02_articles/pkg/errors"
	"github.com/gMerl1on/parsers_articles/02_articles/pkg/jwt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo         repository.StorageUser
	logger       *zap.Logger
	tokenManager jwt.TokenManager
	redisUser    repository.RedisStorageUser
}

type ServiceUser interface {
	CreateUser(ctx context.Context, name, surname, email, password, repeatPassword string, age int) (int, error)
}

func NewUserService(repo repository.StorageUser, log *zap.Logger, tokenManager jwt.TokenManager, redisUser repository.RedisStorageUser) *UserService {
	return &UserService{
		repo:         repo,
		logger:       log,
		tokenManager: tokenManager,
		redisUser:    redisUser,
	}
}

func (u *UserService) CreateUser(ctx context.Context, name, surname, email, password, repeatPassword string, age int) (int, error) {

	if password != repeatPassword {
		u.logger.Warn("Password and repeated password does not match", zap.String("Password", password), zap.String("Repeated password", repeatPassword))
		return 0, er.PasswordRepeatedPassword
	}

	hashedPassword, err := generatePasswordHash(password)
	if err != nil {
		return 0, err
	}

	userID, err := u.repo.CreateUser(ctx, name, surname, email, hashedPassword, age)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (u *UserService) LoginUser(ctx context.Context, email, password string) (*domain.UserByEmail, error) {

	userByEmail, err := u.repo.GetUserByEmail(ctx, email, password)
	if err != nil {
		return nil, err
	}

	if err := checkPassword(password, userByEmail.Password); err != nil {
		return nil, err
	}

}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password due to error %w", err)
	}
	return string(hash), nil
}

func checkPassword(passwordLogin, passwordDB string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordLogin), []byte(passwordDB))
	if err != nil {
		return fmt.Errorf("password does not match")
	}
	return nil
}

func (s *UserService) createSession(ctx context.Context, userUUID string) (jwt.Tokens, error) {

	var (
		res jwt.Tokens
		err error
	)

	res.AccessToken, err = s.tokenManager.NewJWT(userUUID)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return res, err
	}
	session := Session{
		UserUUID:  userUUID,
		ExpiresAt: time.Duration(s.refreshTokenTTL) * time.Minute,
	}

	err = s.redis.SetSession(ctx, res.RefreshToken, session)

	return res, err
}
