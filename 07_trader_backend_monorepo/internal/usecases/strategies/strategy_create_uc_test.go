package strategies

import (
	"context"
	"errors"
	"testing"
	"trader-backend_monorepo/internal/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

		ctx := context.Background()

		stratigiesRepo := mocks.NewStrategiesRepositoryMock(tt)

		stratigiesRepo.
			On("Insert", ctx, mock.Anything).
			Return("", errors.New("some db error when inserting strategy")).
			Once()

		uc := &strategyCreateUC{
			strategiesRepo: stratigiesRepo,
		}

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

		ctx := context.Background()

		stratigiesRepo := mocks.NewStrategiesRepositoryMock(tt)

		stratigiesRepo.
			On("Insert", ctx, mock.Anything).
			Return("strategy-id", nil).
			Once()

		uc := &strategyCreateUC{
			strategiesRepo: stratigiesRepo,
		}

		req := &CreateStrategyRequest{
			Name: "strategy name",
		}

		response, err := uc.Handle(ctx, req)

		assert.Nil(tt, err)
		assert.NotNil(tt, response)
		assert.EqualValues(tt, "strategy-id", response.StrategyID)
	})
}
