package api

import (
	"fmt"
	"net/http"

	"github.com/cjfinnell/go-server/router"
)

func handleEcho(getParams router.GetParamsFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := getParams(r)
		fmt.Fprintf(w, "%s\n", params["input"])
	}
}
