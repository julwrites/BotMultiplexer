package secrets

import (
	"io/ioutil"
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

func GenerateTestData() {
	var secrets SecretsData
	secrets.ADMIN_ID = "123456789"
	secrets.TELEGRAM_ID = "Telegram ID"
	secrets.PROJECT_ID = "Project ID"
	secrets.BUCKET_ID = "Bucket ID"

	data, _ := yaml.Marshal(secrets)

	ioutil.WriteFile("test.yaml", data, os.ModePerm)
}

func TestLoadSecrets(t *testing.T) {
	GenerateTestData()

	data, err := LoadSecrets("test.yaml")
	if err != nil {
		t.Errorf("Failed to load secrets")
	}

	if data.PROJECT_ID != "Project ID" {
		t.Errorf("Failed to load Project ID")
	}

	if data.ADMIN_ID != "123456789" {
		t.Errorf("Failed to load Admin ID")
	}

	if data.TELEGRAM_ID != "Telegram ID" {
		t.Errorf("Failed to load Telegram ID")
	}

	if data.BUCKET_ID != "Bucket ID" {
		t.Errorf("Failed to load Bucket ID")
	}
}
