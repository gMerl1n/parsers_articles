package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gMerl1on/parsers_articles/02_articles/pkg/errors"
	"go.uber.org/zap"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var u UserRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		h.logger.Warn("Failed to decode request body user", zap.Error(err))
		errors.SendHttpError(w, errors.InternalServerError)
		return
	}

	userID, err := h.services.ServiceUser.CreateUser(r.Context(), u.Name, u.Surname, u.Email, u.Passwrod, u.RepeatPassword, u.Age)
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
