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

	fmt.Printf("Conectando ao banco %s:%s (user=%s)\n",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User)
}
