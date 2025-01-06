package parserhb

import (
	"context"

	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/entities"
	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/services"
	"go.uber.org/zap"
)

type RunnerHB struct {
	providerSign   string
	ctx            context.Context
	logger         *zap.Logger
	servCategories services.ServiceCategory
	servArticles   services.ServiceArticle
}

func NewRunnerHB(providerSign string, ctx context.Context, logger *zap.Logger, servCat services.ServiceCategory, servArt services.ServiceArticle) *RunnerHB {
	return &RunnerHB{
		providerSign:   providerSign,
		ctx:            ctx,
		logger:         logger,
		servCategories: servCat,
		servArticles:   servArt,
	}
}

func (r *RunnerHB) RunParserHB() {

	parserHabr := NewParserHabr(r.logger)

	categories, err := r.servCategories.GetCategoriesBySign(r.ctx, r.providerSign)
	if err != nil {
		r.logger.Warn("Failed to get categories", zap.Error(err))
	}

	for _, cat := range categories {

		data := entities.NewDataForParsing(cat.URL, cat.ProviderSign, 123123123123)

		parsedData, _ := parserHabr.ParseLoop(data)

		r.servArticles.CreateArticles(r.ctx, parsedData.Articles)
	}
}
