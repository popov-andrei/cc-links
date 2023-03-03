package server

import (
	"context"
	"log"
	"myapp/internal/handler"
	"myapp/storage"
	"net/http"
)

const (
	port = ":8080"
)

type HTTPServer struct {
	host *http.Server
}

func NewHTTPServer(storage storage.Storage) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.GetHandler(storage)(w, r)
		} else if r.Method == http.MethodPost {
			handler.PostHandler(storage)(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	return &Server{
		server: &http.Server{
			Addr:    port,
			Handler: mux,
		},
	}
}

func (s *HTTPServer) RunServer() {
	log.Println("Starting server...")
	if err := s.host.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not start server: %v\n", err)
	}
}

func (s *HTTPServer) StopServer() {
	log.Println("Stopping server...")
	if err := s.host.Shutdown(context.Background()); err != nil {
		log.Fatalf("Could not stop server: %v\n", err)
	}
}
