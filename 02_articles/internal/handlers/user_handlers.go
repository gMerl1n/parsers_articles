package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gMerl1on/parsers_articles/02_articles/pkg/errors"
	"go.uber.org/zap"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var userRequest UserRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		h.logger.Warn("Failed to decode request body category", zap.Error(err))
		errors.SendHttpError(w, errors.InternalServerError)
		return
	}

	userID, err := h.services.ServiceUser.CreateUser(r.Context(), userRequest.Name, userRequest.Surname, userRequest.Email, userRequest.Age)
	if err != nil {
		h.logger.Warn("Failed to register new user", zap.Error(err))
		errors.SendHttpError(w, errors.InternalServerError)
		return
	}

	marshalledUserID, err := json.Marshal(userID)
	if err != nil {
		h.logger.Warn("Не получилось make marshal", zap.Error(err))
		errors.SendHttpError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledUserID)

}
