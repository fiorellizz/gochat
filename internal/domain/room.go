package domain

import "time"

// Room representa uma sala de conversa
type Room struct {
	ID        string    // UUID
	Name      string
	IsPrivate bool
	OwnerID   string    // User ID
	CreatedAt time.Time
	UpdatedAt time.Time
}
