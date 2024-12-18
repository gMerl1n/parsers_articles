package handlers

import (
	"github.com/gMerl1n/parsers_articles/internal/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(serv *service.Services) *Handler {
	return &Handler{services: serv}
}
