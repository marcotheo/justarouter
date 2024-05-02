package justarouter

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type ServerRouter struct {
	Mux *http.ServeMux
}

type SubRouter struct {
	BasePath string
	Mux      *http.ServeMux
}

type PathParams struct {
	incomingPathReq []string
	pathParams      *map[string]int
}

func setupDynamicPath(pattern string) (string, map[string]int) {
	isDynamic := strings.IndexAny(pattern, ":")

	if isDynamic == -1 {
		return pattern, nil
	}

	path := strings.Split(pattern, ":")[0]
	crumbs := strings.Split(strings.Trim(pattern, "/"), "/")
	pathParamsMap := make(map[string]int)

	for idx, val := range crumbs {
		if val != "" {
			if string(val[0]) == ":" {
				pathParamsMap[val[1:]] = idx
			}
		}
	}

	return path, pathParamsMap
}

func handleRequest(w http.ResponseWriter, r *http.Request, method string, pattern string, handler func(http.ResponseWriter, *http.Request, PathParams), pathParamsMap *map[string]int) {
	var ps PathParams

	if pathParamsMap != nil {
		incomingPathReq := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		ps = PathParams{
			incomingPathReq: incomingPathReq,
			pathParams:      pathParamsMap,
		}
	}

	if r.Method == method {
		handler(w, r, ps)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%v %v api does not exist", r.Method, pattern)
	}
}

func (params *PathParams) Get(key string) (string, error) {
	if (len(params.incomingPathReq)-1) < (*params.pathParams)[key] || (*params.pathParams)[key] == 0 {
		return "", errors.New("path param does not exist")
	}

	return params.incomingPathReq[(*params.pathParams)[key]], nil
}

func CreateRouter() ServerRouter {
	return ServerRouter{
		Mux: http.NewServeMux(),
	}
}

func (server *ServerRouter) POST(pattern string, handler func(http.ResponseWriter, *http.Request, PathParams)) {
	path, pathParamsMap := setupDynamicPath(pattern)
	server.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, "POST", path, handler, &pathParamsMap)
	})
}

func (server *ServerRouter) GET(pattern string, handler func(http.ResponseWriter, *http.Request, PathParams)) {
	path, pathParamsMap := setupDynamicPath(pattern)
	server.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, "GET", path, handler, &pathParamsMap)
	})
}

func (server *ServerRouter) PUT(pattern string, handler func(http.ResponseWriter, *http.Request, PathParams)) {
	path, pathParamsMap := setupDynamicPath(pattern)
	server.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, "PUT", path, handler, &pathParamsMap)
	})
}

func (server *ServerRouter) PATCH(pattern string, handler func(http.ResponseWriter, *http.Request, PathParams)) {
	path, pathParamsMap := setupDynamicPath(pattern)
	server.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, "PATCH", path, handler, &pathParamsMap)
	})
}

func (server *ServerRouter) DELETE(pattern string, handler func(http.ResponseWriter, *http.Request, PathParams)) {
	path, pathParamsMap := setupDynamicPath(pattern)
	server.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, "DELETE", path, handler, &pathParamsMap)
	})
}

func (server *ServerRouter) AddSubRoutes(basePath string, handler func(SubRouter)) {
	subRouter := SubRouter{
		BasePath: basePath,
		Mux:      server.Mux,
	}

	handler(subRouter)
}

func (subRouter *SubRouter) POST(pattern string, handler func(http.ResponseWriter, *http.Request, PathParams)) {
	path, pathParamsMap := setupDynamicPath(subRouter.BasePath + pattern)
	subRouter.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, "POST", path, handler, &pathParamsMap)
	})
}

func (subRouter *SubRouter) GET(pattern string, handler func(http.ResponseWriter, *http.Request, PathParams)) {
	path, pathParamsMap := setupDynamicPath(subRouter.BasePath + pattern)
	subRouter.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, "GET", path, handler, &pathParamsMap)
	})
}

func (subRouter *SubRouter) PUT(pattern string, handler func(http.ResponseWriter, *http.Request, PathParams)) {
	path, pathParamsMap := setupDynamicPath(subRouter.BasePath + pattern)
	subRouter.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, "PUT", path, handler, &pathParamsMap)
	})
}

func (subRouter *SubRouter) PATCH(pattern string, handler func(http.ResponseWriter, *http.Request, PathParams)) {
	path, pathParamsMap := setupDynamicPath(subRouter.BasePath + pattern)
	subRouter.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, "PATCH", path, handler, &pathParamsMap)
	})
}

func (subRouter *SubRouter) DELETE(pattern string, handler func(http.ResponseWriter, *http.Request, PathParams)) {
	path, pathParamsMap := setupDynamicPath(subRouter.BasePath + pattern)
	subRouter.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, "DELETE", path, handler, &pathParamsMap)
	})
}
