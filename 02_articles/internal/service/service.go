package service

import (
	"github.com/gMerl1on/parsers_articles/02_articles/internal/repository"
	"go.uber.org/zap"
)

type Services struct {
	ServiceArticles ServiceArticles
	ServiceCategory ServiceCategory
}

func NewServices(repo *repository.Repositories, logger *zap.Logger) *Services {
	return &Services{
		ServiceArticles: NewArticlesSerivce(repo.Articles, logger),
		ServiceCategory: NewCategoryService(repo.Categories, logger),
	}
}
