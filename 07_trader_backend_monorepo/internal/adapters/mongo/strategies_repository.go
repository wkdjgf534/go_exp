package mongo

import (
	"context"
	"trader-backend_monorepo/internal/domain"
	"trader-backend_monorepo/internal/ports"
)

type StrategiesRepo struct {
}

func NewStrategiesRepository() ports.StrategiesRepository {
	repo := &StrategiesRepo{}

	return repo
}

func (repo *StrategiesRepo) Insert(ctx context.Context, strategy *domain.Strategy) (string, error) {
	// TODO: Insert strategy in MongoDB
	panic("impelemnt me")
}

func (repo *StrategiesRepo) GetByID(ctx context.Context, strategy *domain.Strategy) (*domain.Strategy, error) {
	// TODO: Get strategy form MongoDB
	panic("implement me")
}
