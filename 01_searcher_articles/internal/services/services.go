package services

import (
	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/repository"
	"go.uber.org/zap"
)

type Services struct {
	ServiceArticle ServiceArticle
}

func NewServices(repo repository.Repositories, log *zap.Logger) *Services {
	return &Services{
		ServiceArticle: NewArticlesSerivce(repo.Articles, log),
	}
}
