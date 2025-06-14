package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/SimonKimDev/CoffeeChat/internal/api/routes"
	"github.com/SimonKimDev/CoffeeChat/internal/application"
	"github.com/SimonKimDev/CoffeeChat/internal/infra/config"
	"github.com/SimonKimDev/CoffeeChat/internal/infra/db"
)

func main() {
	var configType string

	cwd, _ := os.Getwd()

	env, _ := os.LookupEnv("APP_ENV")
	if env == "prod" {
		configType = "prod.yaml"
	} else {
		configType = "dev.yaml"
	}

	configPath := filepath.Join(cwd, "configs", configType)
	settings, err := config.Load(configPath)

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	db.InitDatabase(ctx, settings)

	rootMux := http.NewServeMux()

	poster := application.NewPostService()

	routes.RegisterPostRoutes(rootMux, poster)

	srv := &http.Server{
		Addr:    ":" + settings.Server.Port,
		Handler: rootMux,
	}

	log.Println("Listening on http://localhost:8080/hello")
	log.Fatal(srv.ListenAndServe())
}
