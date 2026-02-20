package config

import (
	"os"
)

type Config struct {
	Port         string
	Host         string
	OpenAIAPIKey string
	OpenAIModel  string
	N8NWebhookURL string
}

func Load() *Config {
	return &Config{
		Port:          getEnv("PORT", "8080"),
		Host:          getEnv("HOST", "0.0.0.0"),
		OpenAIAPIKey:  getEnv("OPENAI_API_KEY", ""),
		OpenAIModel:   getEnv("OPENAI_MODEL", "gpt-3.5-turbo"),
		N8NWebhookURL: getEnv("N8N_WEBHOOK_URL", "http://n8n:5678/webhook/code-review"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
