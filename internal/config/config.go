package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DatabaseURL   string
	OpenAIAPIKey  string
	OllamaBaseURL string
	OllamaModel   string
}

func Load() Config {
	// Try to load from .secrets.yaml first, then fall back to environment variables
	if cfg, err := loadFromSecretsFile(".secrets.yaml"); err == nil {
		return cfg
	}

	// Fall back to environment variables
	return Config{
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		OpenAIAPIKey:  os.Getenv("OPENAI_API_KEY"),
		OllamaBaseURL: os.Getenv("OLLAMA_BASE_URL"),
		OllamaModel:   os.Getenv("OLLAMA_MODEL"),
	}
}

func loadFromSecretsFile(filePath string) (Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var secrets map[string]string
	if err := yaml.Unmarshal(data, &secrets); err != nil {
		return Config{}, err
	}

	return Config{
		DatabaseURL:   secrets["database_url"],
		OpenAIAPIKey:  secrets["openai_api_key"],
		OllamaBaseURL: secrets["ollama_base_url"],
		OllamaModel:   secrets["ollama_model"],
	}, nil
}
