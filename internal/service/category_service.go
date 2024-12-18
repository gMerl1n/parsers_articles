package service

import (
	"context"

	"github.com/gMerl1n/parsers_articles/internal/repository"
)

type ServiceCategory interface {
	CreateCategory(ctx context.Context, providerSign string, url string) (int, error)
	GetCategory(ctx context.Context) ([]string, error)
}

type CategoryService struct {
	repo repository.StorageCategory
}

func NewCategoryService(repo repository.StorageCategory) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, providerSign string, url string) (int, error) {
	return 0, nil
}

func (s *CategoryService) GetCategory(ctx context.Context) ([]string, error) {
	return nil, nil
}
