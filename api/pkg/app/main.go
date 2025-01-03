package app

import (
	"akshidas/e-com/pkg/services"
	"log"
	"net/http"
)

type Server struct {
	port       string
	router     *http.ServeMux
	services   *services.Service
	middleware *MiddleWares
}

func (s *Server) Run() {
	log.Printf("🚀 Server started on port %s", s.port)
	log.Fatal(http.ListenAndServe(s.port, s.router))
}

func (s *Server) RegisterRoutes(handler HandleFunc) {
	for path, handler := range handler(s) {
		s.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			handler(w, r)
		})
	}
}

func New(port string, service *services.Service) *Server {
	server := new(Server)
	server.router = http.NewServeMux()
	server.services = service
	server.middleware = newMiddleWare(service.User)

	server.router.HandleFunc("OPTIONS /", func(w http.ResponseWriter, r *http.Request) {
		Cors(w)
	})

	return server
}
