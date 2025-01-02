package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gMerl1on/parsers_articles/02_articles/pkg/errors"
	"go.uber.org/zap"
)

func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var categoryRequest CategoryRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		h.logger.Warn("Failed to decode request body category", zap.Error(err))
		errors.SendHttpError(w, errors.InternalServerError)
		return
	}

	categoryID, err := h.services.ServiceCategory.CreateCategory(r.Context(), categoryRequest.Name, categoryRequest.URL, categoryRequest.ProviderSign)
	if err != nil {
		h.logger.Warn("Не удалось отправить ссылку на категорию в сервисы", zap.Error(err))
		errors.SendHttpError(w, err)
		return
	}

	marshalledCategoryID, err := json.Marshal(categoryID)
	if err != nil {
		h.logger.Warn("Не получилось make marshal", zap.Error(err))
		errors.SendHttpError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledCategoryID)
}

func (h *Handler) GetCategories(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	categories, err := h.services.ServiceCategory.GetCategory(r.Context())
	if err != nil {
		h.logger.Warn("Failed to decode request body category", zap.Error(err))
		errors.SendHttpError(w, errors.InternalServerError)
		return
	}

	marshalledCategories, err := json.Marshal(categories)
	if err != nil {
		h.logger.Warn("Не получилось make marshal categories", zap.Error(err))
		errors.SendHttpError(w, errors.InternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledCategories)
}
