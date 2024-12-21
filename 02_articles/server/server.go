package server

import (
	"context"
	"net/http"

	"github.com/gMerl1on/parsers_articles/02_articles/configs"
	"github.com/gMerl1on/parsers_articles/02_articles/internal/handlers"
	"github.com/gMerl1on/parsers_articles/02_articles/internal/repository"
	"github.com/gMerl1on/parsers_articles/02_articles/internal/service"
	"github.com/gMerl1on/parsers_articles/02_articles/pkg/db"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewHttpServer(ctx context.Context, log *zap.Logger, postgres configs.ConfigPostgres, BindAddr string) (*http.Server, error) {

	db, err := db.NewPostgresDB(ctx, postgres)
	if err != nil {
		log.Fatal("Failed to initialize DB")
	}

	repo := repository.NewRepositories(db, log)
	serv := service.NewServices(repo, log)
	h := handlers.NewHandler(serv, log)

	router := mux.NewRouter()

	// test router
	router.HandleFunc("/api/testOk", h.TestHandler).Methods("GET")

	// handlers categories
	router.HandleFunc("/api/create_category", h.CreateCategory).Methods("POST")
	router.HandleFunc("/api/category", h.GetCategories).Methods("GET")

	return &http.Server{
		Addr:    BindAddr,
		Handler: router,
	}, nil

}
