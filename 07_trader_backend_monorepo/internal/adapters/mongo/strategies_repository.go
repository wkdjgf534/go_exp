package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"trader-backend_monorepo/internal/domain"
	"trader-backend_monorepo/internal/ports"
)

type StrategiesRepo struct {
	strategiesColl *mongo.Collection
}

func NewStrategiesRepository(strategiesColl *mongo.Collection) ports.StrategiesRepository {
	repo := &StrategiesRepo{
		strategiesColl: strategiesColl,
	}

	return repo
}

func (repo *StrategiesRepo) Insert(ctx context.Context, strategy *domain.Strategy) (string, error) {
	// traslate domain.Strategy into StrategyDTO
	strategyDTO, err := fromStrategyCoreToDTO(strategy)
	if err != nil {
		return "", err
	}

	result, err := repo.strategiesColl.InsertOne(ctx, strategyDTO)
	if err != nil {
		return "", fmt.Errorf("error inserting strategy: %w", err)
	}

	strategyID, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return "", fmt.Errorf("error getting inserted id from mongo")
	}

	return strategyID.Hex(), nil
}

func (repo *StrategiesRepo) GetByID(ctx context.Context, strategy *domain.Strategy) (*domain.Strategy, error) {

}
