package service

import (
	"context"

	"github.com/gMerl1on/parsers_articles/02_articles/internal/domain"
	"github.com/gMerl1on/parsers_articles/02_articles/internal/repository"
	"go.uber.org/zap"
)

type ServiceArticles interface {
	GetArticleByID(ctx context.Context, ID int) (domain.Article, error)
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

func (s *ArticlesService) GetArticleByID(ctx context.Context, ID int) (domain.Article, error) {
	return domain.Article{
		Title:   "title",
		Author:  "author",
		Body:    "body",
		Created: 11111.1,
	}, nil
}

func (s *ArticlesService) GetArticles(ctx context.Context) ([]domain.Article, error) {
	return nil, nil
}
