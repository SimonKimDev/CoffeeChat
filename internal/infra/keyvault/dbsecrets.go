package keyvault

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"github.com/SimonKimDev/CoffeeChat/internal/domain/ports"
)

type dbSecretStore struct {
	client              *azsecrets.Client
	connectionSecretKey string
	nameSecretKey       string
}

var version string = ""

func NewDBSecretStore(client *azsecrets.Client, connSecret, nameSecret string) ports.DBSecretProvider {
	return &dbSecretStore{client, connSecret, nameSecret}
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
