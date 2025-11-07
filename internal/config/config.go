package config

import (
	"fmt"
)

// Config contém todas as variáveis de configuração carregadas do ambiente
type Config struct {
	AppEnv   string
	AppPort  string
	LogLevel string

	DB struct {
		Driver   string
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		SSLMode  string
	}

	Redis struct {
		Addr string
	}

	Kafka struct {
		Broker string
	}

	JWT struct {
		Secret string
	}
}

// LoadConfig inicializa todas as configurações do sistema
func LoadConfig() (*Config, error) {
	LoadEnv() // carrega o .env

	cfg := &Config{}

	// App
	cfg.AppEnv = GetEnv("APP_ENV", "development")
	cfg.AppPort = GetEnv("APP_PORT", "8080")
	cfg.LogLevel = GetEnv("LOG_LEVEL", "debug")

	// Database
	cfg.DB.Driver = GetEnv("DB_DRIVER", "postgres")
	cfg.DB.Host = GetEnv("DB_HOST", "localhost")
	cfg.DB.Port = GetEnv("DB_PORT", "5432")
	cfg.DB.User = GetEnv("DB_USER", "")
	cfg.DB.Password = GetEnv("DB_PASSWORD", "")
	cfg.DB.Name = GetEnv("DB_NAME", "")
	cfg.DB.SSLMode = GetEnv("DB_SSLMODE", "disable")

	// Redis / Kafka / JWT
	cfg.Redis.Addr = GetEnv("REDIS_ADDR", "")
	cfg.Kafka.Broker = GetEnv("KAFKA_BROKER", "")
	cfg.JWT.Secret = GetEnv("JWT_SECRET", "")

	// Validação mínima
	required := map[string]string{
		"DB_USER":    cfg.DB.User,
		"DB_PASSWORD": cfg.DB.Password,
		"DB_NAME":    cfg.DB.Name,
		"JWT_SECRET": cfg.JWT.Secret,
	}

	for key, val := range required {
		if val == "" {
			return nil, fmt.Errorf("variável obrigatória ausente: %s", key)
		}
	}

	return cfg, nil
}
