package repository

import (
	"testing"

	"github.com/didiegovieira/go-benchmark-api/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

var clientOpts = mtest.NewOptions().ClientType(mtest.Mock)

func TestSave(t *testing.T) {
	mt := mtest.New(t, clientOpts)

	mt.Run(t.Name(), func(mt *mtest.T) {
		repository := NewBenchmarkMongodb(mt.Client)

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{}},
		})

		err := repository.Save(&entity.Benchmark{})

		assert.Nil(t, err)
	})
}

func TestSaveWithError(t *testing.T) {
	mt := mtest.New(t, clientOpts)

	mt.Run(t.Name(), func(mt *mtest.T) {
		repository := NewBenchmarkMongodb(mt.Client)

		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "mongodb error",
		}))

		err := repository.Save(&entity.Benchmark{})

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "mongodb error")
	})
}
