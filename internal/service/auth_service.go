package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/fiorellizz/gochat/internal/domain"
	"github.com/fiorellizz/gochat/internal/repository"
)

// AuthService provê registro/login com bcrypt + emissão de JWT
type AuthService struct {
	users      repository.UserRepository
	jwtSecret  string
	tokenTTL   time.Duration
}

// NewAuthService cria instância de AuthService
func NewAuthService(users repository.UserRepository, jwtSecret string, tokenTTL time.Duration) *AuthService {
	return &AuthService{
		users:     users,
		jwtSecret: jwtSecret,
		tokenTTL:  tokenTTL,
	}
}

// Register cria usuário (gera UUID, calcula hash da senha)
func (s *AuthService) Register(ctx context.Context, u *domain.User, plainPassword string) (*domain.User, error) {
	if u == nil {
		return nil, errors.New("usuário é nulo")
	}
	if plainPassword == "" {
		return nil, errors.New("a senha é obrigatória")
	}

	// gerar ID (UUID)
	u.ID = uuid.NewString()
	now := time.Now().UTC()
	u.CreatedAt = now
	u.UpdatedAt = now

	// hash da senha
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.PasswordHash = string(hash)

	// persistir via repositório
	if err := s.users.Create(u); err != nil {
		return nil, err
	}

	return u, nil
}

// Login verifica credenciais e retorna JWT (string) + usuário
func (s *AuthService) Login(ctx context.Context, email, plainPassword string) (string, *domain.User, error) {
	u, err := s.users.GetByEmail(email)
	if err != nil {
		return "", nil, err
	}
	if u == nil {
		return "", nil, errors.New("credenciais inválidas")
	}

	// comparar hash
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(plainPassword)); err != nil {
		return "", nil, errors.New("credenciais inválidas")
	}

	// gerar token JWT
	now := time.Now().UTC()
	exp := now.Add(s.tokenTTL)

	claims := jwt.RegisteredClaims{
		Subject:   u.ID,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(exp),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", nil, err
	}

	return signed, u, nil
}
