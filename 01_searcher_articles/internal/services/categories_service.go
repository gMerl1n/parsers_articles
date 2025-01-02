package services

import (
	"context"

	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/entities"
	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/repository"
	"go.uber.org/zap"
)

type CategoryService struct {
	repo   repository.RepoCategory
	logger *zap.Logger
}

func NewCategoryService(repo repository.RepoCategory, log *zap.Logger) *CategoryService {
	return &CategoryService{
		repo:   repo,
		logger: log,
	}
}

type ServiceCategory interface {
	GetCategoriesBySign(ctx context.Context, sign string) ([]entities.Category, error)
}

func (c *CategoryService) GetCategoriesBySign(ctx context.Context, sign string) ([]entities.Category, error) {
	categories, err := c.repo.GetCategoriesBySign(ctx, sign)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
