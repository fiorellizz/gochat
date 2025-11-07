package main

import (
	"github.com/fiorellizz/gochat/internal/config"
	"github.com/fiorellizz/gochat/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log := logger.New("error")
		log.Error("erro ao carregar config: %v", err)
		return
	}

	log := logger.New(cfg.LogLevel)
	log.Info("Configuração carregada com sucesso.")

	db, err := config.ConnectDB(cfg)
	if err != nil {
		log.Error("erro ao conectar no banco: %v", err)
		return
	}
	defer db.Close()

	log.Info("Servidor iniciado com sucesso na porta %s", cfg.AppPort)
}
