package strategies

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"trader-backend_monorepo/internal/domain"
	"trader-backend_monorepo/internal/mocks"
)

func TestStrategyGetByIDUC_Handle(t *testing.T) {
	t.Parallel()

	t.Run("error getting strategy by id returns error", func(tt *testing.T) {
		tt.Parallel()

		ctx := context.Background()

		strategiesRepo := mocks.NewStrategiesRepositoryMock(tt)
		strategiesRepo.On("GetByID", ctx, "strategy-id").
			Return(nil, errors.New("some error from mongodb")).
			Once()

		uc := NewGetByIDUC(strategiesRepo)

		req := &GetStrategyByIDRequest{
			StrategyID: "strategy-id",
		}

		response, err := uc.Handle(ctx, req)

		assert.Nil(tt, response)
		assert.NotNil(tt, err)
		assert.EqualValues(tt, "some error from mongodb", err.Error())
	})

	t.Run("strategy successfully obtained", func(tt *testing.T) {
		tt.Parallel()

		ctx := context.Background()

		mockedStrategy := &domain.Strategy{
			ID:   "strategy-id",
			Name: "name",
		}

		strategiesRepo := mocks.NewStrategiesRepositoryMock(tt)
		strategiesRepo.On("GetByID", ctx, "strategy-id").
			Return(mockedStrategy, nil).
			Once()

		uc := NewGetByIDUC(strategiesRepo)

		req := &GetStrategyByIDRequest{
			StrategyID: "strategy-id",
		}

		response, err := uc.Handle(ctx, req)

		assert.Nil(tt, err)
		assert.NotNil(tt, response)
		assert.NotNil(tt, response.Strategy)
		assert.EqualValues(tt, "strategy-id", response.Strategy.ID)
		assert.EqualValues(tt, "name", response.Strategy.Name)
		assert.EqualValues(tt, "", response.Strategy.Description)
	})
}
