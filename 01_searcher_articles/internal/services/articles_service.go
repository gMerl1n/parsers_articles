package services

import (
	"context"

	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/entities"
	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/repository"
	"go.uber.org/zap"
)

type ArticleService struct {
	repo   repository.RepoArticle
	logger *zap.Logger
}

func NewArticlesSerivce(repo repository.RepoArticle, log *zap.Logger) *ArticleService {
	return &ArticleService{
		repo:   repo,
		logger: log,
	}
}

type ServiceArticle interface {
	CreateArticle(ctx context.Context, article *entities.Article, catID int) (bool, error)
	CreateArticles(ctx context.Context, data *entities.DataForParsing) (bool, error)
}

func (a *ArticleService) CreateArticle(ctx context.Context, article *entities.Article, catID int) (bool, error) {
	res, err := a.repo.CreateArticle(ctx, article, catID)
	if err != nil {
		return false, err
	}

	return res, err

}

func (a *ArticleService) CreateArticles(ctx context.Context, data *entities.DataForParsing) (bool, error) {
	res, err := a.repo.CreateArticles(ctx, data)
	if err != nil {
		return false, err
	}

	return res, err
}
