package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gMerl1on/parsers_articles/02_articles/configs"
	"github.com/gMerl1on/parsers_articles/02_articles/pkg/logging"
	"github.com/gMerl1on/parsers_articles/02_articles/server"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {

	if err := RunServer(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)

}

func RunServer() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	logger, err := logging.InitLogger()
	if err != nil {
		log.Fatal("Failed to load logger")
	}

	// Initialize configs
	config := configs.NewConfig()

	// Initialize server, db, routing
	ctx := context.Background()
	srv, err := server.NewHttpServer(ctx, config.Postgres, config.Bindaddr)

	if err != nil {
		logger.Fatal("Failed to create HTTP server", zap.Error(err))
		return err
	}

	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err = srv.Shutdown(ctx); err != nil {
			fmt.Println("HTTP Server Shutdown")
		}
		close(stopped)
	}()

	logger.Info("Starting API Server...")

	if err = srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	<-stopped

	return nil
}
