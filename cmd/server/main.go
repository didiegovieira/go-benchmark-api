package main

import "github.com/didiegovieira/go-benchmark-api/internal/infrastructure/api"

func main() {
	di, err := InitializeDependencyContainer()
	if err != nil {
		panic(err)
	}

	http := api.NewServer("3000")
	http.RegisterRoutes(di.Routes)
	http.Start()
}
