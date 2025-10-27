package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func NewServer() *http.Server {
	httpRouter := mux.NewRouter()

	httpServer := &http.Server{
		Handler:      httpRouter,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return httpServer
}
