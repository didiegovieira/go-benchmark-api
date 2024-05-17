package main

import "github.com/didiegovieira/go-benchmark-api/di"

func main() {
	api, cleanup, err := di.InitializeApi()
	if err != nil {
		panic("couldn't init api: " + err.Error())
	}

	api.Start()
	defer cleanup()
}
