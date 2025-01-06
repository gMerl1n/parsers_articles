package handlers

import (
	"context"
	"net/http"

	"github.com/gMerl1on/parsers_articles/02_articles/constants"
)

type SortProviderSignOption struct {
	providerSign string
}

func SortProviderSignMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

		providerSign := request.URL.Query().Get("provider_sign")

		if providerSign == "" {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte("Нельзя так делать: не передан знак провайдера"))
			return
		}

		options := SortProviderSignOption{
			providerSign: providerSign,
		}

		ctx := context.WithValue(request.Context(), constants.OptionsContextKey, options)
		request = request.WithContext(ctx)

		h(response, request)
	}
}
