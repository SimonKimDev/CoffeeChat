package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"github.com/SimonKimDev/CoffeeChat/internal/domain"
	_ "github.com/lib/pq"
)

// TODO: Avoid global pointer to DB instance. This should be properly dependency injected
var DB *sql.DB
var version string = ""

type dbSecretStore struct {
	client              *azsecrets.Client
	connectionSecretKey string
	nameSecretKey       string
	userSecretKey       string
}

func InitDatabase(ctx context.Context, settings *domain.Config) {
	var conString, driverName string
	var err error

	cred, err := CreateTokenCredential(settings)

	kvClient, err := azsecrets.NewClient(settings.Azure.KeyVaultUrl, cred, nil)
	if err != nil {
		log.Fatal(err)
	}

	accessToken, err := GetAccessToken(ctx, cred, settings.Database.AadScope)

	if err != nil {
		log.Fatal(err.Error())
	}

	secretStore := dbSecretStore{
		kvClient,
		settings.Database.UrlKeyVaultKey,
		settings.Database.NameKeyVaultKey,
		settings.Database.UserKeyVaultKey,
	}

	serverUrl, err := secretStore.ServerUrl(ctx)
	if err != nil {
		log.Fatal("Error: could not retrieve server url", err.Error())
	}

	dbName, err := secretStore.DBName(ctx)
	if err != nil {
		log.Fatal("Error: could not retrieve db name", err.Error())
	}

	dbUser, err := secretStore.DBUser(ctx)
	if err != nil {
		log.Fatal("Error: could not retrieve db user", err.Error())
	}

	conString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", serverUrl, settings.Database.Ports, dbUser, accessToken.Token, dbName)
	driverName = settings.Database.Driver

	DB, err = sql.Open(driverName, conString)

	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	err = DB.PingContext(ctx)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("CoffeeChat DB Connected Successfully")
}

func (s *dbSecretStore) ServerUrl(ctx context.Context) (string, error) {
	resp, err := s.client.GetSecret(ctx, s.connectionSecretKey, version, nil)
	if err != nil {
		return "", err
	}

	return *resp.Value, nil
}

func (s *dbSecretStore) DBName(ctx context.Context) (string, error) {
	resp, err := s.client.GetSecret(ctx, s.nameSecretKey, version, nil)
	if err != nil {
		return "", err
	}

	return *resp.Value, nil
}

func (s *dbSecretStore) DBUser(ctx context.Context) (string, error) {
	resp, err := s.client.GetSecret(ctx, s.userSecretKey, version, nil)
	if err != nil {
		return "", err
	}

	return *resp.Value, nil
}

func CreateTokenCredential(settings *domain.Config) (azcore.TokenCredential, error) {
	switch settings.Env {
	case "dev":
		return azidentity.NewDefaultAzureCredential(nil)
	case "prod":
		return azidentity.NewManagedIdentityCredential(&azidentity.ManagedIdentityCredentialOptions{
			ID: azidentity.ClientID(settings.Azure.TenantId),
		})
	default:
		return azidentity.NewDefaultAzureCredential(nil)
	}
}

func GetAccessToken(ctx context.Context, cred azcore.TokenCredential, scope string) (azcore.AccessToken, error) {
	options := policy.TokenRequestOptions{Scopes: []string{scope}}
	return cred.GetToken(ctx, options)
}
