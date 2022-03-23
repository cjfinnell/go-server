package server

import (
	"net/http"

	"github.com/cjfinnell/go-server/api"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Run() error {
	api.RegisterRoutes()
	return http.ListenAndServe(":8080", nil)
}
