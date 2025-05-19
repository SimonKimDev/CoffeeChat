package domain

type Config struct {
	Env    string `yaml:"environment"`
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Driver          string `yaml:"driver"`
		UrlKeyVaultKey  string `yaml:"db_url_secret_key"`
		NameKeyVaultKey string `yaml:"db_name_secret_key"`
		UserKeyVaultKey string `yaml:"db_user_secret_key"`
		AadScope        string `yaml:"aad_scope"`
		Ports           int64  `yaml:"ports"`
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
