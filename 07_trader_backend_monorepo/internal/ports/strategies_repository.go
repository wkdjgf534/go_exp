package ports

import (
	"context"

	"trader-backend_monorepo/internal/domain"
)

type StrategiesRepository interface {
	Insert(ctx context.Context, strategy *domain.Strategy) (string, error)
	GetByID(ctx context.Context, strategy *domain.Strategy) (*domain.Strategy, error)
}
