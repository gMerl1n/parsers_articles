package service

import (
	"context"
	"fmt"

	"github.com/gMerl1on/parsers_articles/02_articles/internal/repository"
	er "github.com/gMerl1on/parsers_articles/02_articles/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo   repository.StorageUser
	logger *zap.Logger
}

type ServiceUser interface {
	CreateUser(ctx context.Context, name, surname, email, password, repeatPassword string, age int) (int, error)
}

func NewUserService(repo repository.StorageUser, log *zap.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: log,
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

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password due to error %w", err)
	}
	return string(hash), nil
}
