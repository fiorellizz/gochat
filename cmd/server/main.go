package main

import (
	"fmt"
	"log"

	"github.com/fiorellizz/gochat/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("erro ao carregar config:", err)
	}

	db, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatal("erro ao conectar no banco:", err)
	}
	defer db.Close()

	fmt.Println("Servidor iniciado com sucesso!")
}
