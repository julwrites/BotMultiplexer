package secrets

type SecretsData struct {
	TELEGRAM_ID string `yaml:"TELEGRAM_ID"`
	ADMIN_ID    string `yaml:"ADMIN_ID"`
	PROJECT_ID  string `yaml:"PROJECT_ID"`
	BUCKET_ID   string `yaml:"BUCKET_ID"`
}
