# justarouter

A simple normal go router in other words its just a router

This router uses `net/http` package. nothing much is different with this router so you can just use this router as how you use `net/http` package. I added atleast 3 features to this package which is useful atleast for me and the only reason why I created this router.

### Features available:

1. Defining routes with the http Method.
2. Can add subrouters if you want to organize your routes.
3. Middlewares
   - can add global middlewares
   - can add sub router specific middlewares
   - can add route specific middlewares

example:

```
package main

import (
	"encoding/json"
	justarouter "example/http-server/internal"
	"fmt"
	"net/http"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	Id   string `json:"id"`
}

func pokemonRoutes(subRouter justarouter.SubRouter) {
	subRouter.POST("/info", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Pokemon Info Yet")
	})

	subRouter.GET("/list", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Pokemon List Yet because I'm tired")
	})
}

func subRouterLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Middleware logic before handler
		log.Println("User Sub Router Logger middleware")

		// Call the original handler
		next.ServeHTTP(w, r)

		// Middleware logic after handler
	})
}

func userRoutes(subRouter justarouter.SubRouter) {
	// sub router specific middleware
	subRouter.Use(subRouterLogMiddleware)

	subRouter.GET("/{userId}", func(w http.ResponseWriter, r *http.Request) {
		val := r.PathValue("userId")

		me := UserInfo{
			Name: "marco",
			Age:  "25",
			Id:   val,
		}

		b, err := json.Marshal(me)

		if err != nil {
			fmt.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(b))
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Middleware logic before handler
		log.Println("Authentication middleware")

		// http.Error(w, "Unauthorized", http.StatusUnauthorized)
		// return

		// Call the original handler
		next.ServeHTTP(w, r)

		// Middleware logic after handler
	})
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Middleware logic before handler
		log.Println("Logger middleware")

		// Call the original handler
		next.ServeHTTP(w, r)

		// Middleware logic after handler
	})
}

func main() {
	router := justarouter.CreateRouter()

	// add global middleware
	router.Use(logMiddleware)

	router.AddSubRoutes("/pokemon", pokemonRoutes)
	router.AddSubRoutes("/user", userRoutes)

	router.POST("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "App is Healthy")
	}, authMiddleware) // route specific middleware

	server := http.Server{
		Addr:    ":8080",
		Handler: router.Mux,
	}

	fmt.Println("Server running at port :8080")

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
		return
	}
}
```

Full example can be referred here(https://github.com/marcotheo/samplegoapis)

### Functions

- justarouter.CreateRouter() = creates your server router
- router. POST,GET,PUT,DELETE,PATCH = the simple http methods to define your route also available on subRouter instance
- router.AddSubRoutes(basePath string, subRouter justarouter.SubRouter) = creates sub routes with a base path and subRouter instance where you define your subRoutes
- router.Use = define middlewares
- subRouter.Use = define middlewares on sub route level

### Reasons why I created this

- I just created this for fun
- Trying my best to avoid frameworks and sticking close to the go standard libraries
- For my use in my journey towards creating go applications.
- :) BYE
