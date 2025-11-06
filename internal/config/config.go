package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config representa todas as variáveis de ambiente necessárias
type Config struct {
	AppEnv     string
	AppPort    string
	LogLevel   string

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

// LoadConfig carrega o .env e retorna uma instância preenchida de Config
func LoadConfig() (*Config, error) {
	// Carrega o .env da raiz do projeto (se existir)
	if err := godotenv.Load(); err != nil {
		log.Println("ERRO: arquivo .env não encontrado, usando variáveis do sistema.")
	}

	cfg := &Config{}

	// Variáveis da aplicação
	cfg.AppEnv = os.Getenv("APP_ENV")
	cfg.AppPort = os.Getenv("APP_PORT")
	cfg.LogLevel = os.Getenv("LOG_LEVEL")

	// Banco de Dados
	cfg.DB.Driver = os.Getenv("DB_DRIVER")
	cfg.DB.Host = os.Getenv("DB_HOST")
	cfg.DB.Port = os.Getenv("DB_PORT")
	cfg.DB.User = os.Getenv("DB_USER")
	cfg.DB.Password = os.Getenv("DB_PASSWORD")
	cfg.DB.Name = os.Getenv("DB_NAME")
	cfg.DB.SSLMode = os.Getenv("DB_SSLMODE")

	// Redis / Kafka / JWT
	cfg.Redis.Addr = os.Getenv("REDIS_ADDR")
	cfg.Kafka.Broker = os.Getenv("KAFKA_BROKER")
	cfg.JWT.Secret = os.Getenv("JWT_SECRET")

	// Validação básica das variáveis essenciais
	required := map[string]string{
		"DB_USER": cfg.DB.User,
		"DB_NAME": cfg.DB.Name,
		"DB_PASSWORD": cfg.DB.Password,
		"APP_PORT": cfg.AppPort,
		"JWT_SECRET": cfg.JWT.Secret,
	}

	for key, val := range required {
		if val == "" {
			return nil, fmt.Errorf("variável obrigatória ausente: %s", key)
		}
	}

	return cfg, nil
}
