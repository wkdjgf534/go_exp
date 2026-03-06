package strategies

import (
	"context"

	"trader-backend_monorepo/internal/domain"
	"trader-backend_monorepo/internal/ports"
)

type GetByIDUC interface {
	Handle(ctx context.Context, req *GetStrategyByIDRequest) (*GetStrategyByIDResponse, error)
}

type strategyGetByIDUC struct {
	strategiesRepo ports.StrategiesRepository
}

func NewGetByIDUC(strategiesRepo ports.StrategiesRepository) GetByIDUC {
	result := &strategyGetByIDUC{
		strategiesRepo: strategiesRepo,
	}

	return result
}

type GetStrategyByIDRequest struct {
	StrategyID string
}

type GetStrategyByIDResponse struct {
	Strategy *domain.Strategy
}

func (uc *strategyGetByIDUC) Handle(ctx context.Context, req *GetStrategyByIDRequest) (*GetStrategyByIDResponse, error) {
	strategy, err := uc.strategiesRepo.GetByID(ctx, req.StrategyID)
	if err != nil {
		return nil, err
	}

	response := &GetStrategyByIDResponse{
		Strategy: strategy,
	}

	return response, nil
}
