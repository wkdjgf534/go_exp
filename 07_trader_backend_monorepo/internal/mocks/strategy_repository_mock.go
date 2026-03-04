package mocks

import (
	"context"
	"trader-backend_monorepo/internal/domain"
)

type StrategiesRepositoryMock struct {
	InsertFn  func(ctx context.Context, strategy *domain.Strategy) (string, error)
	GetByIDFn func(ctx context.Context, id string) (*domain.Strategy, error)
}

func (m *StrategiesRepositoryMock) Insert(ctx context.Context, strategy *domain.Strategy) (string, error) {
	return m.InsertFn(ctx, strategy)
}

func (m *StrategiesRepositoryMock) GetByID(ctx context.Context, id string) (*domain.Strategy, error) {
	return m.GetByIDFn(ctx, id)
}
