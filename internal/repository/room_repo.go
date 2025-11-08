package repository

import "github.com/fiorellizz/gochat/internal/domain"

type RoomRepository interface {
	Create(r *domain.Room) error
	GetByID(id string) (*domain.Room, error)
	ListByOwner(ownerID string, limit, offset int) ([]*domain.Room, error)
	Update(r *domain.Room) error
	Delete(id string) error
}
