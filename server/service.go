package server

import (
	"net/http"

	"github.com/cjfinnell/go-server/api"
	"github.com/cjfinnell/go-server/router"
)

type Service struct {
	router router.Router
}

func NewService() *Service {
	return &Service{
		router: router.NewHttpRouter(),
	}
}

func (s *Service) Run() error {
	api.RegisterRoutes(s.router)
	return http.ListenAndServe(":8080", s.router)
}
