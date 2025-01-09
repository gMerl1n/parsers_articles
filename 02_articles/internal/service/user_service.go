package service

import (
	"context"

	"github.com/gMerl1on/parsers_articles/02_articles/internal/repository"
	"go.uber.org/zap"
)

type UserService struct {
	repo   repository.StorageUser
	logger *zap.Logger
}

type ServiceUser interface {
	CreateUser(ctx context.Context, name, surname, email string, age int) (int, error)
}

func NewUserService(repo repository.StorageUser, log *zap.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: log,
	}
}

func (u *UserService) CreateUser(ctx context.Context, name, surname, email string, age int) (int, error) {
	userID, err := u.repo.CreateUser(ctx, name, surname, email, age)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
