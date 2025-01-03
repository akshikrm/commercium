package app

import (
	"akshidas/e-com/pkg/services"
	"log"
	"net/http"
)

type Server struct {
	port     string
	router   *http.ServeMux
	services *services.Service
}

func (s *Server) Run() {
	log.Printf("🚀 Server started on port %s", s.port)
	log.Fatal(http.ListenAndServe(s.port, s.router))
}

func New(port string, service *services.Service) *Server {
	server := new(Server)
	server.router = http.NewServeMux()
	server.services = service

	server.router.HandleFunc("OPTIONS /", func(w http.ResponseWriter, r *http.Request) {
		Cors(w)
	})

	server.registerUserRoutes()
	server.registerProductRoutes()
	server.registerProductCategoryRoutes()
	server.registerCartRoutes()
	server.registerPurchaseRoutes()
	return server
}
