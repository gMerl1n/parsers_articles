package server

import (
	"net/http"
)

func NewHttpServer(BindAddr string) (*http.Server, error) {

	return &http.Server{
		Addr: BindAddr,
	}, nil

}
