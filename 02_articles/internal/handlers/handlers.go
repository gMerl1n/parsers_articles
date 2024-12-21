package handlers

import (
	"github.com/gMerl1on/parsers_articles/02_articles/internal/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(serv *service.Services) *Handler {
	return &Handler{services: serv}
}
