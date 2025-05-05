package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/SimonKimDev/CoffeeChat/internal/domain"
	"github.com/SimonKimDev/CoffeeChat/internal/domain/ports"
	_ "github.com/microsoft/go-mssqldb/azuread"
)

var db *sql.DB

func InitDatabase(ctx context.Context, cred ports.TokenProvider, settings *domain.Config, secretStore ports.DBSecretProvider) {
	var conString, driverName string
	var err error

	accessToken, err := cred.GetToken(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	switch settings.Env {
	case "dev":
		conString = settings.Database.Driver
		driverName = settings.Database.Driver
	case "prod":
		serverUrl, err := secretStore.ServerUrl(ctx)
		if err != nil {
			log.Fatalf("Error: could not retrieve server url", err.Error())
		}
		dbName, err := secretStore.DBName(ctx)
		if err != nil {
			log.Fatalf("Error: could not retrieve db name", err.Error())
		}
		conString = fmt.Sprintf("server=%s;database=%s;fedauth=ActiveDirectoryDefault;Token=%s", serverUrl, dbName, accessToken.Token)
		driverName = settings.Database.Driver
	}

	db, err = sql.Open(driverName, conString)

	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	err = db.PingContext(ctx)

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected to DB")
}
