package main

import (
	adapter "github.com/SimonKimDev/CoffeeChat/internal/adapter/http"
	"github.com/SimonKimDev/CoffeeChat/internal/application"
	"log"
	"net/http"
)

func main() {
	rootMux := http.NewServeMux()

	greeter := application.NewGreeterService()
	adapter.RegisterRoutes(rootMux, greeter)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: rootMux,
	}

	log.Println("Listening on http://localhost:8080/hello")
	log.Fatal(srv.ListenAndServe())
}
