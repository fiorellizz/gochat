package repository

import "github.com/fiorellizz/gochat/internal/domain"

type UserRepository interface {
	Create(u *domain.User) error
	GetByID(id string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Update(u *domain.User) error
	Delete(id string) error
}
