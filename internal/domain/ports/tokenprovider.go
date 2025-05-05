package ports

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

type TokenProvider interface {
	GetToken(ctx context.Context) (azcore.AccessToken, error)
	GetCred() azcore.TokenCredential
}
