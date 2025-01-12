package service

import (
	"github.com/gMerl1on/parsers_articles/02_articles/internal/repository"
	"github.com/gMerl1on/parsers_articles/02_articles/pkg/jwt"
	"go.uber.org/zap"
)

type Services struct {
	ServiceArticles ServiceArticles
	ServiceCategory ServiceCategory
	ServiceUser     ServiceUser
	TokenManager    jwt.TokenManager
}

func NewServices(repo *repository.Repositories, logger *zap.Logger, TokenManager jwt.TokenManager) *Services {
	return &Services{
		ServiceArticles: NewArticlesSerivce(repo.Articles, logger),
		ServiceCategory: NewCategoryService(repo.Categories, logger),
		ServiceUser:     NewUserService(repo.Users, logger, TokenManager, repo.UserRedis),
	}
}
