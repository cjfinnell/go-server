package router

import "net/http"

type Params map[string]string

type GetParamsFunc func(req *http.Request) Params

type Router interface {
	HandlerFunc(method, path string, handle http.HandlerFunc)
	GetParams(req *http.Request) Params
	http.Handler
}
