package infrastructure

import (
	"context"
	"imp-backend/domain"
)

// Auth Repository Interface
type AuthRepository interface {
	GetUser(ctx context.Context, params *domain.Users) (*domain.Users, error)
	CreateUser(ctx context.Context, params *domain.Users) (int64, error)
	UpdateUser(ctx context.Context, params *domain.Users) error
	SelectUser(ctx context.Context, params *domain.UserParams) ([]domain.Users, error)
}
