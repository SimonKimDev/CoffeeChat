package domain

type Config struct {
	Env    string `yaml:"environment"`
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Type         string `yaml:"type"`
		DatabaseUrl  string `yaml:"db_url"`
		DatabaseName string `yaml:"db_name"`
		Scope        string `yaml:"scope"`
	} `yaml:"database"`
	Auth struct {
		JwtSecret          string `yaml:"jwt_secret"`
		TokenExpiryMinutes int32  `yaml:"token_expiry_minutes"`
	}
	Azure struct {
		KeyVaultUrl string `yaml:"keyvaulturl"`
		TenantId    string `envconfig:"AZURE_TENANT_ID"`
	}
}
