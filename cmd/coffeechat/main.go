package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	adapter "github.com/SimonKimDev/CoffeeChat/internal/adapter/http"
	"github.com/SimonKimDev/CoffeeChat/internal/application"
	"github.com/SimonKimDev/CoffeeChat/internal/infra/auth"
	"github.com/SimonKimDev/CoffeeChat/internal/infra/config"
	"github.com/SimonKimDev/CoffeeChat/internal/infra/db"
	"github.com/SimonKimDev/CoffeeChat/internal/infra/keyvault"
)

func main() {
	cwd, _ := os.Getwd()
	configPath := filepath.Join(cwd, "configs", "prod.yaml")
	settings, err := config.Load(configPath)

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	tp, err := auth.NewTokenProvider(settings)
	if err != nil {
		log.Fatal(err)
	}

	kvClient, err := azsecrets.NewClient(settings.Azure.KeyVaultUrl, tp.GetCred(), nil)
	if err != nil {
		log.Fatal(err)
	}

	dbSecrets := keyvault.NewDBSecretStore(kvClient, settings.Database.UrlKeyVaultKey, settings.Database.NameKeyVaultKey)

	db.InitDatabase(ctx, tp, settings, dbSecrets)

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
