package server

import (
	"context"
	"net/http"

	"github.com/gMerl1on/parsers_articles/02_articles/configs"
	"github.com/gMerl1on/parsers_articles/02_articles/internal/handlers"
	"github.com/gMerl1on/parsers_articles/02_articles/internal/repository"
	"github.com/gMerl1on/parsers_articles/02_articles/internal/service"
	"github.com/gMerl1on/parsers_articles/02_articles/pkg/db/postgres_storage"
	"github.com/gMerl1on/parsers_articles/02_articles/pkg/db/redis_storage"
	"github.com/gMerl1on/parsers_articles/02_articles/pkg/jwt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewHttpServer(ctx context.Context, log *zap.Logger, postgres configs.ConfigPostgres, redisConf configs.ConfigRedis, BindAddr string, tokenManager jwt.TokenManager) (*http.Server, error) {

	db, err := postgres_storage.NewPostgresDB(ctx, postgres)
	if err != nil {
		log.Fatal("Failed to initialize DB")
	}

	redisDB, err := redis_storage.NewRedisClient(redisConf)
	if err != nil {
		log.Fatal("Failed to initialize Redis")
	}

	repo := repository.NewRepositories(db, redisDB, log)
	serv := service.NewServices(repo, log, tokenManager)
	h := handlers.NewHandler(serv, log)

	router := mux.NewRouter()

	// test router
	router.HandleFunc("/api/testOk", h.TestHandler).Methods("GET")

	// handlers categories
	router.HandleFunc("/api/create_category", h.CreateCategory).Methods("POST")
	router.HandleFunc("/api/category", h.GetCategories).Methods("GET")

	// handlers articles
	router.HandleFunc("/api/articles", h.GetArticles).Methods("GET")
	router.HandleFunc("/api/articlesBySign", handlers.SortProviderSignMiddleware(h.GetArticlesBySign)).Methods("GET")
	router.HandleFunc("/api/articlesByCategory", handlers.CategoryIDMiddleware(h.GetArticlesByCategory)).Methods("GET")

	// handlers users
	router.HandleFunc("/api/createUser", h.CreateUser).Methods("POST")
	router.HandleFunc("/api/loginUser", h.LoginUser).Methods("GET")

	return &http.Server{
		Addr:    BindAddr,
		Handler: router,
	}, nil

}
