package service

import (
	"context"

	"github.com/gMerl1on/parsers_articles/02_articles/internal/domain"
	"github.com/gMerl1on/parsers_articles/02_articles/internal/repository"
	"go.uber.org/zap"
)

type ServiceArticles interface {
	GetArticles(ctx context.Context) ([]domain.Article, error)
}

type ArticlesService struct {
	repo   repository.StorageArticles
	logger *zap.Logger
}

func NewArticlesSerivce(repo repository.StorageArticles, log *zap.Logger) *ArticlesService {
	return &ArticlesService{
		repo:   repo,
		logger: log,
	}
}

func (s *ArticlesService) GetArticles(ctx context.Context) ([]domain.Article, error) {

	articles, err := s.repo.GetArticles(ctx)
	if err != nil {
		return nil, err
	}

	return articles, err
}
