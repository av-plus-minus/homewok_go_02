package server

import (
	"net/http"

	"gateway/internal/controller"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	pingController *controller.PingController
	router         *chi.Mux
}

func NewServer(pingController *controller.PingController) *Server {
	s := &Server{
		pingController: pingController,
		router:         chi.NewRouter(),
	}
	s.configureRouter()
	return s
}

func (s *Server) configureRouter() {
	s.router.Get("/ping", s.pingController.Ping)
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
