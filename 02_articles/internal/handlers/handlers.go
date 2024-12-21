package handlers

import (
	"github.com/gMerl1on/parsers_articles/02_articles/internal/service"
	"go.uber.org/zap"
)

type Handler struct {
	services *service.Services
	logger   *zap.Logger
}

func NewHandler(serv *service.Services, log *zap.Logger) *Handler {
	return &Handler{services: serv, logger: log}
}
