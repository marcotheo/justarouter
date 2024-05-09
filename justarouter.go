package justarouter

import (
	"net/http"
)

type ServerRouter struct {
	Mux *http.ServeMux
}

type SubRouter struct {
	BasePath string
	Mux      *http.ServeMux
}

func CreateRouter() ServerRouter {
	return ServerRouter{
		Mux: http.NewServeMux(),
	}
}

func (server *ServerRouter) POST(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	server.Mux.HandleFunc("POST " + pattern, handler)
}

func (server *ServerRouter) GET(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	server.Mux.HandleFunc("GET " + pattern, handler)
}

func (server *ServerRouter) PUT(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	server.Mux.HandleFunc("PUT " + pattern, handler)
}

func (server *ServerRouter) PATCH(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	server.Mux.HandleFunc("PATCH " + pattern, handler)
}

func (server *ServerRouter) DELETE(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	server.Mux.HandleFunc("DELETE " + pattern, handler)
}

func (server *ServerRouter) AddSubRoutes(basePath string, handler func(SubRouter)) {
	subRouter := SubRouter{
		BasePath: basePath,
		Mux:      server.Mux,
	}

	handler(subRouter)
}

func (subRouter *SubRouter) POST(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	subRouter.Mux.HandleFunc("POST " + subRouter.BasePath + pattern, handler)
}

func (subRouter *SubRouter) GET(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	subRouter.Mux.HandleFunc("GET " + subRouter.BasePath + pattern, handler)
}

func (subRouter *SubRouter) PUT(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	subRouter.Mux.HandleFunc("PUT " + subRouter.BasePath + pattern, handler)
}

func (subRouter *SubRouter) PATCH(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	subRouter.Mux.HandleFunc("PATCH " + subRouter.BasePath + pattern, handler)
}

func (subRouter *SubRouter) DELETE(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	subRouter.Mux.HandleFunc("DELETE " + subRouter.BasePath + pattern, handler)
}
