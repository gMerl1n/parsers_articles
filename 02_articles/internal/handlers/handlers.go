package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (h *Handler) TestHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	ok, err := json.Marshal("ok")
	if err != nil {
		fmt.Println("Не получилось make marshal categories")
	}

	w.Write(ok)

}
