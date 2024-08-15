package justarouter

import "net/http"

func reverseSlice(slice []Middleware) []Middleware {
	left := 0
	right := len(slice) - 1
	for left < right {
		slice[left], slice[right] = slice[right], slice[left]
		left++
		right--
	}
	return slice
}

// Function that accepts multiple middleware functions
func applyMiddlewares(handler http.Handler, globalMiddlewares []Middleware, middlewares ...Middleware) http.Handler {
	mergedMiddlewares := append(globalMiddlewares, middlewares...)
	reverseMiddlewares := reverseSlice(mergedMiddlewares)

	for _, middleware := range reverseMiddlewares {
		handler = middleware(handler)
	}

	return handler
}
