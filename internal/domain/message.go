package domain

import "time"

// Message representa uma mensagem enviada em uma sala
type Message struct {
	ID          string    // UUID
	RoomID      string
	SenderID    string
	Content     string
	ContentType string    // ex: "text", "image", "system"
	Delivered   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
