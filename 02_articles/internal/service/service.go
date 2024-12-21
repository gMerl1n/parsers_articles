package service

import "github.com/gMerl1n/parsers_articles/internal/repository"

type Services struct {
	ServiceArticles ServiceArticles
	ServiceCategory ServiceCategory
}

func NewServices(repo *repository.Repositories) *Services {
	return &Services{
		ServiceArticles: NewArticlesSerivce(repo.Articles),
		ServiceCategory: NewCategoryService(repo.Categories),
	}
}
