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
		return nil, fmt.Errorf("invalid input strategy")
	}

	result := &StrategyDTO{
		Name:        input.Name,
		Description: input.Description,
	}

	if input.ID != "" {
		id, err := bson.ObjectIDFromHex(input.ID)
		if err != nil {
			return nil, fmt.Errorf("invalid strategy id: '%s'", input.ID)
		}

		result.ID = id
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
