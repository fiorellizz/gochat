package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/fiorellizz/gochat/internal/domain"
	"github.com/fiorellizz/gochat/internal/repository"
)

// MessagePublisher é uma interface opcional para publicar mensagens em fila/hub.
// Pode ficar nil se você ainda não implementou fila/hub.
type MessagePublisher interface {
	Publish(ctx context.Context, m *domain.Message) error
}

// ChatService gerencia envio e leitura de mensagens
type ChatService struct {
	messages  repository.MessageRepository
	publisher MessagePublisher // opcional, pode ser nil
}

// NewChatService cria ChatService
func NewChatService(messages repository.MessageRepository, publisher MessagePublisher) *ChatService {
	return &ChatService{
		messages:  messages,
		publisher: publisher,
	}
}

// Send envia (persiste) uma mensagem e publica se houver publisher
func (s *ChatService) Send(ctx context.Context, m *domain.Message) (*domain.Message, error) {
	if m == nil {
		return nil, errorsNew("message is nil")
	}
	// gerar id e timestamps se não existirem
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	now := time.Now().UTC()
	if m.CreatedAt.IsZero() {
		m.CreatedAt = now
	}
	m.UpdatedAt = now

	// Persistir
	if err := s.messages.Create(m); err != nil {
		return nil, err
	}

	// Publicar (opcional)
	if s.publisher != nil {
		// publicar de forma não bloqueante poderia ser uma decisão; aqui chamamos diretamente
		if err := s.publisher.Publish(ctx, m); err != nil {
			// se publicar falhar, decidimos não falhar o envio; só logar/retornar o erro conforme necessidade.
			// Aqui retornamos o erro para que o chamador possa decidir.
			return m, err
		}
	}

	return m, nil
}

// ListByRoom lista mensagens de uma sala
func (s *ChatService) ListByRoom(ctx context.Context, roomID string, limit, offset int) ([]*domain.Message, error) {
	return s.messages.ListByRoom(roomID, limit, offset)
}
