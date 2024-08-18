package justarouter

import (
	"net/http"
	"strings"
	"time"
)

type ServerRouter struct {
	Mux         *http.ServeMux
	Handler     *http.Handler
	Middlewares []Middleware
}

type SubRouter struct {
	BasePath    string
	Mux         *http.ServeMux
	Middlewares []Middleware
}

type Middleware func(http.Handler) http.Handler

// CorsOptions holds configuration for CORS
type CorsOptions struct {
	AllowedOrigins   []string      // List of allowed origins
	AllowedMethods   []string      // List of allowed HTTP methods
	AllowCredentials bool          // Whether credentials are allowed
	AllowedHeaders   []string      // List of allowed headers
	MaxAge           time.Duration // Optional: Max age of the preflight response
}

type ServerRouterOptions struct {
	CORS CorsOptions
}

func corsMiddleware(next http.Handler, corsOptions CorsOptions) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		for _, allowedOrigin := range corsOptions.AllowedOrigins {
			if origin == allowedOrigin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		if len(corsOptions.AllowedHeaders) > 0 {
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(corsOptions.AllowedHeaders, ", "))
		}

		if len(corsOptions.AllowedMethods) > 0 {
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(corsOptions.AllowedMethods, ", "))
		}

		if corsOptions.AllowCredentials {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CreateRouter(options ServerRouterOptions) ServerRouter {
	muxInstance := http.NewServeMux()
	newHandler := corsMiddleware(muxInstance, options.CORS)

	return ServerRouter{
		Mux:         muxInstance,
		Handler:     &newHandler,
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
