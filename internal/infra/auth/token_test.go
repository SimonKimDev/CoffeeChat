package auth

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/SimonKimDev/CoffeeChat/internal/domain"
)

// Helper for Testing Purposes
type fakeCred struct {
	token string
}

func (f fakeCred) GetToken(context.Context, policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return azcore.AccessToken{
		Token:     f.token,
		ExpiresOn: time.Now().Add(10 * time.Minute),
	}, nil
}

func TestBuild_ChoosesExpectedCredential(t *testing.T) {
	cfgCallback := func(env string) *domain.Config {
		return &domain.Config{
			Env: env,
		}
	}

	env := "test"
	expectedType := reflect.TypeOf(&azidentity.DefaultAzureCredential{})

	cred, err := Build(cfgCallback(env))

	if err != nil {
		t.Fatalf("Build(%q) returned error: %v", env, err)
	}

	if reflect.TypeOf(cred) != expectedType {
		t.Errorf("Build(%q) returned %T, want %v", env, cred, expectedType)
	}
}

func TestTokenProviderAdapter_GetToken(t *testing.T) {
	const dummy = "dummyToken"
	tp := tokenProviderAdapter{
		Cred:     fakeCred{token: dummy},
		AadScope: "https://example/.default",
	}

	accessToken, err := tp.GetToken(context.Background())

	if err != nil {
		t.Fatalf("GetToken returned error: %v", err)
	}

	if accessToken.Token != dummy {
		t.Errorf("GetToken token = %q, want %q", accessToken.Token, dummy)
	}
}
