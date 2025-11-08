package domain

import "time"

// User representa um usuário do sistema (domínio puro)
type User struct {
	ID           string    // UUID (gerado pela camada de infra)
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
