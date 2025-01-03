package app

import (
	"akshidas/e-com/pkg/services"
	"log"
	"net/http"
)

type Server struct {
	port   string
	router *http.ServeMux
}

func (s *Server) Run() {
	log.Printf("🚀 Server started on port %s", s.port)
	log.Fatal(http.ListenAndServe(s.port, s.router))
}

func New(port string, service *services.Service) *Server {
	server := new(Server)
	server.router = http.NewServeMux()
	server.registerRoutes(service)
	return server
}
