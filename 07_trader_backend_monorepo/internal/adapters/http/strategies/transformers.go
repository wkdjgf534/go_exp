package strategies

import "trader-backend_monorepo/internal/domain"

func fromStrategyCoreToHTTP(input *domain.Strategy) *Strategy {
	if input == nil {
		return nil
	}

	result := &Strategy{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
	}

	return result
}
