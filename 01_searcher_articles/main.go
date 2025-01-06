package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gMerl1on/parsers_articles/01_searcher_articles/configs"
	parserhb "github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/parsers/parser_hb"
	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/repository"
	"github.com/gMerl1on/parsers_articles/01_searcher_articles/internal/services"
	"github.com/gMerl1on/parsers_articles/01_searcher_articles/pkg/db"
	"github.com/gMerl1on/parsers_articles/01_searcher_articles/pkg/logging"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Searcher articles service")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	runnerHB := parserhb.NewRunnerHB("HB", ctx, logger, serv.ServiceCategory, serv.ServiceArticle)
	runnerHB.RunParserHB()

}
