package stratigies

import (
	"context"
	"strings"

	"trader-backend_monorepo/internal/domain"
	"trader-backend_monorepo/internal/ports"
	"trader-backend_monorepo/pkg/apierrors"
)

type StrategyCreateUC interface {
	Handle(ctx context.Context, req *CreateStrategyRequest) (*CreateStrategyResponse, error)
}

type strategyCreateUC struct {
	strategiesRepo ports.StrategiesRepository
}

type CreateStrategyRequest struct {
	Name         string
	Descriiption string
}

func NewStrategyCreateUC(strategiesRepo ports.StrategiesRepository) StrategyCreateUC {
	result := &strategyCreateUC{
		strategiesRepo: strategiesRepo,
	}

	return result
}

type CreateStrategyResponse struct {
	StrategyID string
}

func (uc *strategyCreateUC) Handle(ctx context.Context, req *CreateStrategyRequest) (*CreateStrategyResponse, error) {
	strategy, err := uc.fromCreateStrategyReqToStrategy(req)
	if err != nil {
		return nil, err
	}

	strategyID, err := uc.strategiesRepo.Insert(ctx, strategy)
	if err != nil {
		return nil, err
	}

	response := &CreateStrategyResponse{
		StrategyID: strategyID,
	}

	return response, nil
}

func (uc *strategyCreateUC) fromCreateStrategyReqToStrategy(req *CreateStrategyRequest) (*domain.Strategy, error) {
	if req == nil {
		return nil, apierrors.NewBadRequestError("invalid request")
	}

	if req.Name = strings.TrimSpace(req.Name); req.Name == "" {
		return nil, apierrors.NewBadRequestError("invalid name")
	}

	result := &domain.Strategy{
		Name:        req.Name,
		Description: req.Descriiption,
	}

	return result, nil
}
