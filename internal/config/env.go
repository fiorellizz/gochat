package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv tenta carregar o arquivo .env e retorna as variáveis de ambiente
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado — usando variáveis do sistema.")
	}
}

// GetEnv recupera variáveis do sistema de forma segura
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
