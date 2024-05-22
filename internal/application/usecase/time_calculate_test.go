package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/didiegovieira/go-benchmark-api/internal/application/dto"
	"github.com/stretchr/testify/assert"
)

func TestTimeCalculateImplementationExecute(t *testing.T) {
	timeCalculate := NewTimeCalculate()

	t.Run("should measure execution time of a function correctly", func(t *testing.T) {
		input := dto.TimeCalculateInput{
			Name: "TestFunc",
			Func: func() {
				time.Sleep(100 * time.Millisecond)
			},
		}

		result, err := timeCalculate.Execute(context.Background(), input)

		assert.NoError(t, err)
		assert.Equal(t, "TestFunc", result.Name)
		assert.True(t, result.Duration >= 100*time.Millisecond.Milliseconds())
	})

	t.Run("should return result entity with correct name and duration", func(t *testing.T) {
		input := dto.TimeCalculateInput{
			Name: "TestFunc",
			Func: func() {
				time.Sleep(50 * time.Millisecond)
			},
		}

		result, err := timeCalculate.Execute(context.Background(), input)

		assert.NoError(t, err)
		assert.Equal(t, "TestFunc", result.Name)
		assert.True(t, result.Duration >= 50*time.Millisecond.Milliseconds())
	})
}
