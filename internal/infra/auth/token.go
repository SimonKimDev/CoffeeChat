package auth

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/SimonKimDev/CoffeeChat/internal/domain"
	"github.com/SimonKimDev/CoffeeChat/internal/domain/ports"
)

type tokenProviderAdapter struct {
	Cred     azcore.TokenCredential
	AadScope string
}

func Build(settings *domain.Config) (azcore.TokenCredential, error) {
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

func (p tokenProviderAdapter) GetToken(ctx context.Context) (azcore.AccessToken, error) {
	options := policy.TokenRequestOptions{Scopes: []string{p.AadScope}}
	return p.Cred.GetToken(ctx, options)
}

func (p tokenProviderAdapter) GetCred() azcore.TokenCredential {
	return p.Cred
}

func NewTokenProvider(settings *domain.Config) (ports.TokenProvider, error) {
	cred, err := Build(settings)
	if err != nil {
		return nil, err
	}

	return tokenProviderAdapter{
		cred,
		settings.Database.AadScope,
	}, nil
}
