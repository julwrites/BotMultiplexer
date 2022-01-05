package secrets

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Secrets handling for bot

func LoadSecrets(filePath string) (SecretsData, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading secrets: %v", err)
		return nil, err
	}

	// log.Printf("Raw Secrets: %s", strings.ReplaceAll(string(data), "\n", ""))

	var env SecretsData
	err = yaml.Unmarshal([]byte(data), &env)
	if err != nil {
		log.Printf("Error unmarshaling secrets: %v", err)
		return nil, err
	}

	return env, nil
}
