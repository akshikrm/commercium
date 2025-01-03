package app

import (
	"akshidas/e-com/pkg/services"
	"log"
	"net/http"
)

type app struct {
	port   string
	router *http.ServeMux
}

func (s *app) Run() {
	log.Printf("🚀 Server started on port %s", s.port)
	log.Fatal(http.ListenAndServe(s.port, s.router))
}

func New(port string, service *services.Service) *app {
	server := new(app)
	server.router = http.NewServeMux()
	server.registerRoutes(service)
	return server
}
