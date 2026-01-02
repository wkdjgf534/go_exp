package mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"trader-backend_monorepo/internal/domain"
	"trader-backend_monorepo/internal/ports"
	"trader-backend_monorepo/pkg/apierrors"
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
		// BadRequest error
		return "", apierrors.NewBadRequestError(err.Error())
	}

	result, err := repo.strategiesColl.InsertOne(ctx, strategyDTO)
	if err != nil {
		// Internal error
		return "", apierrors.NewInternalServerError("error creating strategy")
	}

	strategyID, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		// Internal error
		return "", apierrors.NewInternalServerError("error getting inserted id")
	}

	return strategyID.Hex(), nil
}

func (repo *StrategiesRepo) GetByID(ctx context.Context, id string) (*domain.Strategy, error) {
	mongoID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, apierrors.NewBadRequestError(fmt.Sprintf("invalid strategy id: '%s'", id))
	}

	filter := bson.M{
		"_id": mongoID,
	}

	result := repo.strategiesColl.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			// NotFound error
			return nil, apierrors.NewNotFoundError(fmt.Sprintf("strategy '%s' not found", id))
		}

		// Internal error
		return nil, apierrors.NewInternalServerError("error getting strategy by id")
	}

	var strategyDTO StrategyDTO
	if err = result.Decode(&strategyDTO); err != nil {
		// Internal error
		return nil, apierrors.NewInternalServerError(fmt.Sprintf("error getting strategy '%s'", id))
	}

	strategy := fromStrategyDTOToCore(strategyDTO)

	return &strategy, nil
}
