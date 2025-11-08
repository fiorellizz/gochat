package service

import (
	"context"

	"github.com/fiorellizz/gochat/internal/domain"
	"github.com/fiorellizz/gochat/internal/repository"
)

// RoomService trata criação/listagem de salas
type RoomService struct {
	rooms repository.RoomRepository
}

// NewRoomService cria RoomService
func NewRoomService(rooms repository.RoomRepository) *RoomService {
	return &RoomService{rooms: rooms}
}

// Create cria uma sala
func (s *RoomService) Create(ctx context.Context, r *domain.Room) error {
	// validações básicas (nome)
	if r == nil {
		return ErrNilRoom
	}
	if r.Name == "" {
		return ErrInvalidRoomName
	}
	return s.rooms.Create(r)
}

// GetByID retorna sala por id
func (s *RoomService) GetByID(ctx context.Context, id string) (*domain.Room, error) {
	return s.rooms.GetByID(id)
}

// ListByOwner lista salas por dono com paginação simples
func (s *RoomService) ListByOwner(ctx context.Context, ownerID string, limit, offset int) ([]*domain.Room, error) {
	return s.rooms.ListByOwner(ownerID, limit, offset)
}

// Erros de domínio simples
var (
	ErrNilRoom         = errorsNew("room is nil")
	ErrInvalidRoomName = errorsNew("room name is empty")
)

// errorsNew é uma pequena função utilitária para evitar import cycles com errors
func errorsNew(msg string) error { return &simpleError{msg} }

type simpleError struct{ s string }
func (e *simpleError) Error() string { return e.s }
