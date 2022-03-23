package api

import (
	"net/http"

	"github.com/cjfinnell/go-server/router"
)

func RegisterRoutes(r router.Router) {
	r.HandlerFunc(http.MethodGet, "/hello", handleHello)
	r.HandlerFunc(http.MethodGet, "/echo/:input", handleEcho(r.GetParams))
}
