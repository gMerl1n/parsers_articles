package service

import (
	"context"

	"github.com/gMerl1n/parsers_articles/internal/repository"
)

type Services interface {
	CreateCategory(ctx context.Context, providerSign string, url string) (int, error)
	GetCategory(ctx context.Context) ([]string, error)
}

type service struct {
	storage repository.Storage
}

func NewServices(repo *repository.Repository) Services {
	return &service{storage: repo}
}

func (s *service) CreateCategory(ctx context.Context, providerSign string, url string) (int, error) {
	categoryID, err := s.storage.CreateCategory(ctx, url, providerSign)
	if err != nil {
		return 0, err
	}

	return categoryID, nil
}

func (s *service) GetCategory(ctx context.Context) ([]string, error) {
	return nil, nil
}
