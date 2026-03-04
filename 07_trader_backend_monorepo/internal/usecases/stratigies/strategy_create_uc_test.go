package stratigies

import (
	"context"
	"errors"
	"testing"
	"trader-backend_monorepo/internal/domain"
	"trader-backend_monorepo/internal/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateStrategyUC_Handle(t *testing.T) {
	t.Parallel()

	t.Run("invalid request returns error", func(tt *testing.T) {
		tt.Parallel()

		uc := &strategyCreateUC{}

		ctx := context.Background()
		req := &CreateStrategyRequest{}

		response, err := uc.Handle(ctx, req)

		assert.Nil(tt, response)
		assert.NotNil(tt, err)
		assert.EqualValues(tt, "invalid name", err.Error())
	})

	t.Run("error creating strategy returns error", func(tt *testing.T) {
		tt.Parallel()

		stratigiesRepo := mocks.StrategiesRepositoryMock{}
		stratigiesRepo.InsertFn = func(ctx context.Context, strategy *domain.Strategy) (string, error) {
			return "", errors.New("some db error when inserting strategy")
		}

		uc := &strategyCreateUC{
			strategiesRepo: &stratigiesRepo,
		}

		ctx := context.Background()
		req := &CreateStrategyRequest{
			Name: "strategy name",
		}

		response, err := uc.Handle(ctx, req)

		assert.Nil(tt, response)
		assert.NotNil(tt, err)
		assert.EqualValues(tt, "some db error when inserting strategy", err.Error())
	})

	t.Run("strategy successfully created", func(tt *testing.T) {
		tt.Parallel()

		stratigiesRepo := mocks.StrategiesRepositoryMock{}
		stratigiesRepo.InsertFn = func(ctx context.Context, strategy *domain.Strategy) (string, error) {
			return "strategy-id", nil
		}

		uc := &strategyCreateUC{
			strategiesRepo: &stratigiesRepo,
		}

		ctx := context.Background()
		req := &CreateStrategyRequest{
			Name: "strategy name",
		}

		response, err := uc.Handle(ctx, req)

		assert.Nil(tt, err)
		assert.NotNil(tt, response)
		assert.EqualValues(tt, "strategy-id", response.StrategyID)
	})
}
