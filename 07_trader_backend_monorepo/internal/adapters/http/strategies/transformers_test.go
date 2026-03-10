package strategies

import (
	"testing"
	"trader-backend_monorepo/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestFromStrategyCoreToHTTP(t *testing.T) {
	t.Parallel()

	t.Run("nil input returns nil", func(tt *testing.T) {
		tt.Parallel()

		output := fromStrategyCoreToHTTP(nil)

		assert.Nil(tt, output)
	})

	t.Run("strategy successfully transformed", func(tt *testing.T) {
		tt.Parallel()

		input := &domain.Strategy{
			ID:          "id",
			Name:        "name",
			Description: "description",
		}
		output := fromStrategyCoreToHTTP(input)

		assert.NotNil(tt, output)
		assert.EqualValues(tt, input.ID, output.ID)
		assert.EqualValues(tt, input.Name, output.Name)
		assert.EqualValues(tt, input.Description, output.Description)
	})
}
