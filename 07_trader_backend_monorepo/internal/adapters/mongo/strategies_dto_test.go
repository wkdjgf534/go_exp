package mongo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"

	"trader-backend_monorepo/internal/domain"
)

func Test_FromStrategyCoreToDTO(t *testing.T) {
	t.Parallel()

	mockedMongoID := "68e6bc21731dcf55202ad7fc"
	mockedObjectID, err := bson.ObjectIDFromHex(mockedMongoID)
	assert.Nil(t, err)

	testCases := []struct {
		Name   string
		Input  *domain.Strategy
		Output *StrategyDTO
		Error  error
	}{
		{
			Name:   "nil input returns error",
			Input:  nil,
			Output: nil,
			Error:  errors.New("invalid input strategy"),
		},
		{
			Name: "empty mongo id should not populate id",
			Input: &domain.Strategy{
				ID:          "",
				Name:        "name",
				Description: "description",
			},
			Output: &StrategyDTO{
				ID:          bson.NilObjectID,
				Name:        "name",
				Description: "description",
			},
			Error: nil,
		},
		{
			Name: "invalid mongo id return error",
			Input: &domain.Strategy{
				ID:          "invalid mongo id",
				Name:        "name",
				Description: "description",
			},
			Output: nil,
			Error:  errors.New("invalid strategy id: 'invalid mongo id'"),
		},
		{
			Name: "strategy successfully processed",
			Input: &domain.Strategy{
				ID:          mockedMongoID,
				Name:        "name",
				Description: "description",
			},
			Output: &StrategyDTO{
				ID:          mockedObjectID,
				Name:        "name",
				Description: "description",
			},
			Error: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()

			output, err := fromStrategyCoreToDTO(tc.Input)

			if tc.Error != nil {
				assert.Nil(tt, output)
				assert.NotNil(tt, err)
				assert.EqualValues(tt, tc.Error.Error(), err.Error())
				return
			}

			assert.Nil(tt, err)
			assert.NotNil(tt, output)
			assert.EqualValues(tt, tc.Output.ID, output.ID)
			assert.EqualValues(tt, tc.Output.Name, output.Name)
			assert.EqualValues(tt, tc.Output.Description, output.Description)
		})
	}
}

func Test_FromStrategyDTOToCore(t *testing.T) {
	t.Parallel()
}
