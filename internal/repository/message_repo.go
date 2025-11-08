package repository

import "github.com/fiorellizz/gochat/internal/domain"

type MessageRepository interface {
	Create(m *domain.Message) error
	GetByID(id string) (*domain.Message, error)
	ListByRoom(roomID string, limit, offset int) ([]*domain.Message, error)
	ListRecentByRoom(roomID string, limit int) ([]*domain.Message, error)
	Update(m *domain.Message) error
	Delete(id string) error
}
