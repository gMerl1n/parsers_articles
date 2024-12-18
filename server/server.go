package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gMerl1n/parsers_articles/configs"
	"github.com/gMerl1n/parsers_articles/internal/handlers"
	"github.com/gMerl1n/parsers_articles/internal/repository"
	"github.com/gMerl1n/parsers_articles/internal/service"
	"github.com/gMerl1n/parsers_articles/pkg/db"
	"github.com/gorilla/mux"
)

func NewHttpServer(ctx context.Context, postgres configs.ConfigPostgres, BindAddr string) (*http.Server, error) {

	db, err := db.NewPostgresDB(ctx, postgres)
	if err != nil {
		fmt.Println("Ошибка инициализации БД")
	}

	repo := repository.NewRepositories(db)
	serv := service.NewServices(repo)
	h := handlers.NewHandler(serv)

	router := mux.NewRouter()

	// handlers categories
	router.HandleFunc("/api/create_cat", h.CreateCategory).Methods("POST")
	router.HandleFunc("/api/cat", h.GetCategories).Methods("GET")

	return &http.Server{
		Addr:    BindAddr,
		Handler: router,
	}, nil

}
