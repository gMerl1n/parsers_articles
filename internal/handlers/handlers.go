package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gMerl1n/parsers_articles/internal/service"
)

type Handler struct {
	services service.Services
}

func NewHandler(serv service.Services) *Handler {
	return &Handler{services: serv}
}

func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var requestProviderUrl RequestProviderUrl
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&requestProviderUrl); err != nil {
		fmt.Println("Ошибка во время парсинга ссылки в RequestUrl")
	}

	categoryID, err := h.services.CreateCategory(r.Context(), requestProviderUrl.Url)
	if err != nil {
		fmt.Println("Не удалось отправить ссылку на категорию в сервисы")
	}

	marshalledCategoryID, err := json.Marshal(categoryID)
	if err != nil {
		fmt.Println("Не получилось make marshal")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledCategoryID)
}

func (h *Handler) GetCategories(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	categories, err := h.services.GetCategory(r.Context())
	if err != nil {
		fmt.Println("Не получилось достать все категории")
	}

	marshalledCategories, err := json.Marshal(categories)
	if err != nil {
		fmt.Println("Не получилось make marshal categories")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledCategories)
}
