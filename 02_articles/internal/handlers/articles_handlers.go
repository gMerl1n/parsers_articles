package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gMerl1on/parsers_articles/02_articles/constants"
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

func (h *Handler) GetArticlesBySign(w http.ResponseWriter, r *http.Request) {

	providerSignOptions := r.Context().Value(constants.OptionsContextKey).(SortProviderSignOption)

	articlesBySign, err := h.services.ServiceArticles.GetArticlesBySign(r.Context(), providerSignOptions.providerSign)
	if err != nil {
		h.logger.Warn("Failed to get articles", zap.Error(err))
		errors.SendHttpError(w, errors.InternalServerError)
	}

	marshalledArticlesBySign, err := json.Marshal(articlesBySign)
	if err != nil {
		h.logger.Warn("Не получилось make marshal articles", zap.Error(err))
		errors.SendHttpError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledArticlesBySign)

}

func (h *Handler) GetArticlesByCategory(w http.ResponseWriter, r *http.Request) {

	categoryID := r.Context().Value(constants.OptionsContextKey).(CategoryIDOption)

	articlesByCategory, err := h.services.ServiceArticles.GetArticlesByCategory(r.Context(), categoryID.ID)
	if err != nil {
		h.logger.Warn("Failed to get articles", zap.Error(err))
		errors.SendHttpError(w, errors.InternalServerError)
	}

	marshalledArticlesByCategory, err := json.Marshal(articlesByCategory)
	if err != nil {
		h.logger.Warn("Не получилось make marshal articles", zap.Error(err))
		errors.SendHttpError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledArticlesByCategory)

}
