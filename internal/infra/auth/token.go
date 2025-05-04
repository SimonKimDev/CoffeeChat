package auth

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

func NewTokenCredential(env, clientId string) (azcore.TokenCredential, error) {
	var cred azcore.TokenCredential
	var err error

	switch env {
	case "dev":
		cred, err = azidentity.NewDefaultAzureCredential(nil)
	case "prod":
		//cred, err = azidentity.NewManagedIdentityCredential(&azidentity.ManagedIdentityCredentialOptions{
		//		ID: azidentity.ClientID(clientId),
		//	})
		cred, err = azidentity.NewDefaultAzureCredential(nil)
	}

	if err != nil {
		return nil, fmt.Errorf("Error: failed to create credential: %v", err)
	}

	return cred, nil
}

func GetAccessToken(ctx context.Context, scope string, cred azcore.TokenCredential) (azcore.AccessToken, error) {
	accessToken, err := cred.GetToken(ctx, policy.TokenRequestOptions{Scopes: []string{scope}})
	if err != nil {
		log.Fatalf("Error: failed to get accessToken", err.Error())
	}

	return accessToken, nil
}
