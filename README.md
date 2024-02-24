# justarouter
A simple normal go router in other words its just a router

This router uses `net/http` package. nothing much is different with this router so you can just use this router as how you use `net/http` package. I added atleast 3 features to this package which is useful atleast for me and the only reason why I created this router.

### Features available:
1. Defining routes with the http Method instead of using switch case.
2. Can add path parameters by using the symbol ":" to define them.
3. Can add subrouters if you want to organize your routes.
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
	subRouter.POST("/info", func(w http.ResponseWriter, r *http.Request, pp justarouter.PathParams) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Pokemon Info Yet")
	})

	subRouter.GET("/list", func(w http.ResponseWriter, r *http.Request, pp justarouter.PathParams) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Pokemon List Yet because I'm tired")
	})
}

func userRoutes(subRouter justarouter.SubRouter) {
	subRouter.GET("/:userId", func(w http.ResponseWriter, r *http.Request, pp justarouter.PathParams) {
		val, err := pp.Get("userId")

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

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

func main() {
	router := justarouter.CreateRouter()

	router.AddSubRoutes("/pokemon", pokemonRoutes)
	router.AddSubRoutes("/user", userRoutes)

	router.POST("/health", func(w http.ResponseWriter, r *http.Request, pp justarouter.PathParams) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "App is Healthy")
	})

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

### Functions
- justarouter.CreateRouter() = creates your server router
- router. POST,GET,PUT,DELETE,PATCH = the simple http methods to define your route also available on subRouter instance
- router.AddSubRoutes(basePath string, subRouter justarouter.SubRouter) = creates sub routes with a base path and subRouter instance where you define your subRoutes

### Path Parameter
When working with path parameters like /user/:userId you can access it through the 3rd property of the function handler `pp justarouter.PathParams` and can access its value with `pp.Get("userId")` it should return a `(string, error)`

### Reasons why I created this
- I just created this for fun
- Trying my best to avoid frameworks and sticking close to the go standard libraries
- For my use in my journey towards creating go applications.
- :) BYE


