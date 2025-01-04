package app

import (
	"akshidas/e-com/pkg/handlers"
	"log"
	"net/http"
)

type Server struct {
	port     string
	router   *http.ServeMux
	handlers *handlers.Handler
}

func (s *Server) Run() {
	log.Printf("🚀 Server started on port %s", s.port)
	log.Fatal(http.ListenAndServe(s.port, s.router))
}

func (s *Server) RegisterRoutes(handler routesFunc) {
	for path, handler := range handler(s) {
		s.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			handler(w, r)
		})
	}
}

func New(port string, handler *handlers.Handler) *Server {
	server := new(Server)
	server.port = port
	server.router = http.NewServeMux()
	server.handlers = handler

	server.router.HandleFunc("OPTIONS /", func(w http.ResponseWriter, r *http.Request) {
		handlers.Cors(w)
	})

	return server
}
