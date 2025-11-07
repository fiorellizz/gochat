package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// ConnectDB cria e testa a conexão com o PostgreSQL
func ConnectDB(cfg *Config) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com o banco: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao conectar no banco: %w", err)
	}

	fmt.Println("✅ Conectado ao banco de dados PostgreSQL com sucesso.")
	return db, nil
}
