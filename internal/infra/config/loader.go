package config

import (
	"fmt"
	"os"

	"github.com/SimonKimDev/CoffeeChat/internal/domain"
	"gopkg.in/yaml.v3"
)

var tenantId string = "AZURE_TENANT_ID"

func Load(path string) (*domain.Config, error) {
	content, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("Error: failed to read config file: %w", err)
	}

	var config domain.Config

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, fmt.Errorf("Error: failed to parse config file: %w", err)
	}

	config.Azure.TenantId, err = getEnv(tenantId)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)

	if value != "" {
		return "", fmt.Errorf("Error: Required Env Var %s is not set", key)
	}

	return value, nil
}
