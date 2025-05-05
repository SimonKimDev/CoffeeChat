package config

import (
	"os"
	"testing"
)

func TestLoadConfig_Success(t *testing.T) {
	expectedEnv := "test"
	expectedHost := "1.2.3.4"
	expectedPort := "567"
	expectedDbType := "sqllite"
	expectedDatabaseUrl := "test_connection_string"
	expectedSecret := "test_secret"
	expectedToken := 5
	expectedKv := "test_keyvault"
	expectedTenantId := "12345"

	os.Setenv(TenantId, expectedTenantId)
	defer os.Unsetenv(TenantId)

	path := "../../../configs/test.yaml"
	cfg, err := Load(path)

	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if cfg.Env != expectedEnv {
		t.Errorf("Expected environment %q, got %q", expectedEnv, cfg.Env)
	}

	if cfg.Server.Host != expectedHost {
		t.Errorf("Expected host %q, got %q", expectedHost, cfg.Server.Host)
	}

	if cfg.Server.Port != expectedPort {
		t.Errorf("Expected port %q, got %q", expectedPort, cfg.Server.Port)
	}

	if cfg.Database.Driver != expectedDbType {
		t.Errorf("Expected DB type %q, got %q", expectedDbType, cfg.Database.Driver)
	}

	if cfg.Database.UrlKeyVaultKey != expectedDatabaseUrl {
		t.Errorf("Expected DB connection string %q, got %q", expectedDatabaseUrl, cfg.Database.UrlKeyVaultKey)
	}

	if cfg.Auth.JwtSecret != expectedSecret {
		t.Errorf("Expected JWT secret %q, got %q", expectedSecret, cfg.Auth.JwtSecret)
	}

	if int(cfg.Auth.TokenExpiryMinutes) != expectedToken {
		t.Errorf("Expected token expiry %d, got %d", expectedToken, cfg.Auth.TokenExpiryMinutes)
	}

	if cfg.Azure.KeyVaultUrl != expectedKv {
		t.Errorf("Expected key vault name %q, got %q", expectedKv, cfg.Azure.KeyVaultUrl)
	}

	if cfg.Azure.TenantId != expectedTenantId {
		t.Errorf("Expected tenant ID %q from env, got %q", expectedTenantId, cfg.Azure.TenantId)
	}
}

func TestLoadConfig_FailsWithoutTenantId(t *testing.T) {
	os.Unsetenv(TenantId)

	path := "../../../configs/test.yaml"
	_, err := Load(path)

	if err == nil {
		t.Errorf("Expected Failure due to missing AZURE_TENANT_ID, but Config Loading Succeeded")
	}
}
