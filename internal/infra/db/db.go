package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/SimonKimDev/CoffeeChat/internal/infra/auth"
	"github.com/microsoft/go-mssqldb/azuread"
)

var db *sql.DB

func InitDatabase(ctx context.Context, env, dbType, dbUrl, dbName, scope string, cred azcore.TokenCredential) {
	var err error
	var conString, driverName string
	accessToken, err := auth.GetAccessToken(ctx, scope, cred)

	if err != nil {
		log.Fatal(err.Error())
	}

	switch env {
	case "dev":
		conString = dbUrl
		driverName = dbType
	case "prod":
		conString = fmt.Sprintf("server=%s;database=%s;fedauth=ActiveDirectoryDefault;Token=%s", dbUrl, dbName, accessToken.Token)
		driverName = azuread.DriverName
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
