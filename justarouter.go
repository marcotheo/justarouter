package justarouter

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type ServerRouter struct {
	Mux         *http.ServeMux
	Middlewares []Middleware
}

type SubRouter struct {
	BasePath    string
	Mux         *http.ServeMux
	Middlewares []Middleware
}

func CreateRouter() ServerRouter {
	return ServerRouter{
		Mux:         http.NewServeMux(),
		Middlewares: []Middleware{},
	}
}

func (server *ServerRouter) Use(middleware Middleware) {
	server.Middlewares = append(server.Middlewares, middleware)
}

func (server *ServerRouter) POST(pattern string, handler func(http.ResponseWriter, *http.Request), middlewares ...Middleware) {
	finalHander := applyMiddlewares(http.HandlerFunc(handler), server.Middlewares, middlewares...)
	server.Mux.Handle("POST "+pattern, finalHander)
}

func (server *ServerRouter) GET(pattern string, handler func(http.ResponseWriter, *http.Request), middlewares ...Middleware) {
	finalHander := applyMiddlewares(http.HandlerFunc(handler), server.Middlewares, middlewares...)
	server.Mux.Handle("GET "+pattern, finalHander)
}

func (server *ServerRouter) PUT(pattern string, handler func(http.ResponseWriter, *http.Request), middlewares ...Middleware) {
	finalHander := applyMiddlewares(http.HandlerFunc(handler), server.Middlewares, middlewares...)
	server.Mux.Handle("PUT "+pattern, finalHander)
}

func (server *ServerRouter) PATCH(pattern string, handler func(http.ResponseWriter, *http.Request), middlewares ...Middleware) {
	finalHander := applyMiddlewares(http.HandlerFunc(handler), server.Middlewares, middlewares...)
	server.Mux.Handle("PATCH "+pattern, finalHander)
}

func (server *ServerRouter) DELETE(pattern string, handler func(http.ResponseWriter, *http.Request), middlewares ...Middleware) {
	finalHander := applyMiddlewares(http.HandlerFunc(handler), server.Middlewares, middlewares...)
	server.Mux.Handle("DELETE "+pattern, finalHander)
}

func (server *ServerRouter) AddSubRoutes(basePath string, handler func(SubRouter)) {
	// copy the global middleware to sub router middlewares
	fromGlobalMiddlewares := make([]Middleware, len(server.Middlewares))
	copy(fromGlobalMiddlewares, server.Middlewares)

	subRouter := SubRouter{
		BasePath:    basePath,
		Mux:         server.Mux,
		Middlewares: fromGlobalMiddlewares,
	}

	handler(subRouter)
}

func (subRouter *SubRouter) Use(middleware Middleware) {
	subRouter.Middlewares = append(subRouter.Middlewares, middleware)
}

func (subRouter *SubRouter) POST(pattern string, handler func(http.ResponseWriter, *http.Request), middlewares ...Middleware) {
	finalHander := applyMiddlewares(http.HandlerFunc(handler), subRouter.Middlewares, middlewares...)
	subRouter.Mux.Handle("POST "+subRouter.BasePath+pattern, finalHander)
}

func (subRouter *SubRouter) GET(pattern string, handler func(http.ResponseWriter, *http.Request), middlewares ...Middleware) {
	finalHander := applyMiddlewares(http.HandlerFunc(handler), subRouter.Middlewares, middlewares...)
	subRouter.Mux.Handle("GET "+subRouter.BasePath+pattern, finalHander)
}

func (subRouter *SubRouter) PUT(pattern string, handler func(http.ResponseWriter, *http.Request), middlewares ...Middleware) {
	finalHander := applyMiddlewares(http.HandlerFunc(handler), subRouter.Middlewares, middlewares...)
	subRouter.Mux.Handle("PUT "+subRouter.BasePath+pattern, finalHander)
}

func (subRouter *SubRouter) PATCH(pattern string, handler func(http.ResponseWriter, *http.Request), middlewares ...Middleware) {
	finalHander := applyMiddlewares(http.HandlerFunc(handler), subRouter.Middlewares, middlewares...)
	subRouter.Mux.Handle("PATCH "+subRouter.BasePath+pattern, finalHander)
}

func (subRouter *SubRouter) DELETE(pattern string, handler func(http.ResponseWriter, *http.Request), middlewares ...Middleware) {
	finalHander := applyMiddlewares(http.HandlerFunc(handler), subRouter.Middlewares, middlewares...)
	subRouter.Mux.Handle("DELETE "+subRouter.BasePath+pattern, finalHander)
}
