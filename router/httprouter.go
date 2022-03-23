package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HttpRouter struct {
	httprouter.Router
}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{*httprouter.New()}
}

func (r *HttpRouter) GetParams(req *http.Request) Params {
	params := make(Params)
	for _, pair := range httprouter.ParamsFromContext(req.Context()) {
		params[pair.Key] = pair.Value
	}
	return params
}
