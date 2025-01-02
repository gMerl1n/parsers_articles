package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gMerl1on/parsers_articles/01_searcher_articles/configs"
	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/entities"
	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/parsers"
	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/repository"
	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/services"
	"github.com/gMerl1on/parsers_articles/01_searcher_articles/pkg/db"
	"github.com/gMerl1on/parsers_articles/01_searcher_articles/pkg/logging"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("Searcher articles service")

	logger, err := logging.InitLogger()
	if err != nil {
		log.Fatal("Failed to load logger")
	}

	// Initialize configs
	config := configs.NewConfig()

	// Initialize server, db, routing
	ctx := context.Background()

	db, err := db.NewPostgresDB(ctx, config.Postgres)
	if err != nil {
		log.Fatal("Failed to initialize DB")
	}

	repo := repository.NewRepositories(db, logger)
	serv := services.NewServices(repo, logger)

	parserHabr := parsers.NewParserHabr(logger)

	categories, err := serv.ServiceCategory.GetCategoriesBySign(ctx, "HB")
	if err != nil {
		logger.Warn("Failed to get categories", zap.Error(err))
	}

	for _, cat := range categories {
		data := entities.NewDataForParsing(cat.URL, cat.ProviderSign, 123123123123)
		parserHabr.ParseLoop(data)

	}

}
