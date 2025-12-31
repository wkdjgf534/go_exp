package mongo

import (
	"fmt"
	"trader-backend_monorepo/internal/domain"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type StrategyDTO struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
}

func fromStrategyCoreToDTO(input *domain.Strategy) (*StrategyDTO, error) {
	if input == nil {
		return nil, fmt.Errorf("invalid imput strategy")
	}

	id, err := bson.ObjectIDFromHex(input.ID) // convert strong to bson.ObjectID
	if err != nil {
		return nil, fmt.Errorf("invalid strategy id: %s", input.ID)
	}

	result := &StrategyDTO{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
	}

	return result, nil
}

func fromStrategyDTOToCore(input StrategyDTO) domain.Strategy {
	result := domain.Strategy{
		ID:          input.ID.Hex(),
		Name:        input.Name,
		Description: input.Description,
	}

	return result
}
