package application

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

const dbUrlKey string = "DBConnectionString"
const dbNameKey string = "dbcoffeechat"

type KeyVaultService struct {
	client *azsecrets.Client
}

func NewKeyVaultService(url string, cred azcore.TokenCredential) (*KeyVaultService, error) {
	client, err := azsecrets.NewClient(url, cred, nil)
	if err != nil {
		return nil, err
	}

	return &KeyVaultService{
		client: client,
	}, nil
}

func (s *KeyVaultService) GetDbSecrets() (string, string, error) {

	version := ""
	dbUrl, err := s.client.GetSecret(context.TODO(), dbUrlKey, version, nil)
	if err != nil {
		return "", "", fmt.Errorf("Error: failed to retrieve db connection string")
	}

	dbName, err := s.client.GetSecret(context.TODO(), dbNameKey, version, nil)

	return *dbUrl.Value, *dbName.Value, nil
}
