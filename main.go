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

	"github.com/gMerl1n/parsers_articles/configs"
	"github.com/gMerl1n/parsers_articles/server"
	"github.com/joho/godotenv"
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

	// Initialize configs
	config := configs.NewConfig()

	// Initialize server, db, routing
	ctx := context.Background()
	srv, err := server.NewHttpServer(ctx, config.Postgres, config.Bindaddr)

	if err != nil {
		fmt.Println("Failed to start server", err.Error())
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

	fmt.Println("starting API Server...")

	if err = srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	<-stopped

	return nil
}
