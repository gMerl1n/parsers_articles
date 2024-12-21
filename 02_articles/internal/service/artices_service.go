package service

import (
	"context"

	"github.com/gMerl1on/parsers_articles/02_articles/internal/repository"
	tt "github.com/gMerl1on/parsers_articles/02_articles/internal/types"
)

type ServiceArticles interface {
	GetArticleByID(ctx context.Context, ID int) (tt.Article, error)
	GetArticles(ctx context.Context) ([]tt.Article, error)
}

type ArticlesService struct {
	repo repository.StorageArticles
}

func NewArticlesSerivce(repo repository.StorageArticles) *ArticlesService {
	return &ArticlesService{
		repo: repo,
	}
}

func (s *ArticlesService) GetArticleByID(ctx context.Context, ID int) (tt.Article, error) {
	return tt.Article{
		Title:   "title",
		Author:  "author",
		Body:    "body",
		Created: 11111.1,
	}, nil
}

func (s *ArticlesService) GetArticles(ctx context.Context) ([]tt.Article, error) {
	return nil, nil
}
