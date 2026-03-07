package strategies

import (
	"context"

	"trader-backend_monorepo/internal/ports"
)

type Service interface {
	Create(ctx context.Context, req *CreateStrategyRequest) (*CreateStrategyResponse, error)
	GetByID(ctx context.Context, req *GetStrategyByIDRequest) (*GetStrategyByIDResponse, error)
}

type service struct {
	createUC  CreateUC
	getByIDUC GetByIDUC
}

func NewService(strategiesRepo ports.StrategiesRepository) Service {
	createUC := NewCreateUC(strategiesRepo)
	getByIDUC := NewGetByIDUC(strategiesRepo)

	svc := &service{
		createUC:  createUC,
		getByIDUC: getByIDUC,
	}

	return svc
}

func (s *service) Create(ctx context.Context, req *CreateStrategyRequest) (*CreateStrategyResponse, error) {
	return s.createUC.Handle(ctx, req)
}

func (s *service) GetByID(ctx context.Context, req *GetStrategyByIDRequest) (*GetStrategyByIDResponse, error) {
	return s.getByIDUC.Handle(ctx, req)
}
