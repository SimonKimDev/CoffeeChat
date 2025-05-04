package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"

	adapter "github.com/SimonKimDev/CoffeeChat/internal/adapter/http"
	"github.com/SimonKimDev/CoffeeChat/internal/application"
	"github.com/SimonKimDev/CoffeeChat/internal/infra/auth"
	"github.com/SimonKimDev/CoffeeChat/internal/infra/config"
	"github.com/SimonKimDev/CoffeeChat/internal/infra/db"
)

func main() {
	cwd, _ := os.Getwd()
	configPath := filepath.Join(cwd, "configs", "prod.yaml")
	settings, err := config.Load(configPath)

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	cred, err := auth.NewTokenCredential(settings.Env, settings.Azure.TenantId)
	kvService, err := application.NewKeyVaultService(settings.Azure.KeyVaultUrl, cred)

	dbUrl, dbName, err := kvService.GetDbSecrets()

	db.InitDatabase(ctx, settings.Env, settings.Database.Type, dbUrl, dbName, settings.Database.Scope, cred)

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
