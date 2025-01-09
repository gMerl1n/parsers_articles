package service

import (
	"github.com/gMerl1on/parsers_articles/02_articles/internal/repository"
	"go.uber.org/zap"
)

type Services struct {
	ServiceArticles ServiceArticles
	ServiceCategory ServiceCategory
	ServiceUser     ServiceUser
}

func NewServices(repo *repository.Repositories, logger *zap.Logger) *Services {
	return &Services{
		ServiceArticles: NewArticlesSerivce(repo.Articles, logger),
		ServiceCategory: NewCategoryService(repo.Categories, logger),
		ServiceUser:     NewUserService(repo.Users, logger),
	}
}
