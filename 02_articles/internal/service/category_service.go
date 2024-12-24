package service

import (
	"context"

	"github.com/gMerl1on/parsers_articles/02_articles/internal/domain"
	"github.com/gMerl1on/parsers_articles/02_articles/internal/repository"
	"go.uber.org/zap"
)

type ServiceCategory interface {
	CreateCategory(ctx context.Context, name string, url string) (int, error)
	GetCategory(ctx context.Context) ([]domain.Category, error)
}

type CategoryService struct {
	repo   repository.StorageCategory
	logger *zap.Logger
}

func NewCategoryService(repo repository.StorageCategory, log *zap.Logger) *CategoryService {
	return &CategoryService{
		repo:   repo,
		logger: log,
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, name string, url string) (int, error) {

	categoryID, err := s.repo.CreateCategory(ctx, name, url)
	if err != nil {
		return 0, err
	}

	return categoryID, nil

}

func (s *CategoryService) GetCategory(ctx context.Context) ([]domain.Category, error) {

	categories, err := s.repo.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
