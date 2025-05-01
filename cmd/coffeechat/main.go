package main

import (
	"net/http"

	"github.com/SimonKimDev/CoffeeChat/api"
)

func main() {
	mux := http.NewServeMux()

	api.RegisterRoutes(mux)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	srv.ListenAndServe()
}
