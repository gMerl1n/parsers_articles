package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gMerl1on/parsers_articles/02_articles/pkg/errors"
	"go.uber.org/zap"
)

func (h *Handler) GetArticles(w http.ResponseWriter, r *http.Request) {

	articles, err := h.services.ServiceArticles.GetArticles(r.Context())
	if err != nil {
		h.logger.Warn("Failed to get articles", zap.Error(err))
		errors.SendHttpError(w, errors.InternalServerError)
	}

	marshalledArticles, err := json.Marshal(articles)
	if err != nil {
		h.logger.Warn("Не получилось make marshal articles", zap.Error(err))
		errors.SendHttpError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledArticles)

}
