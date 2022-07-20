package repo_test

import (
	"log"
	"testing"

	"github.com/covid-19-project/api/src/pkg/repo"
	"github.com/covid-19-project/api/src/pkg/services"
	"github.com/go-playground/assert/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestGetNewCasesPerLocation(t *testing.T) {
	databaseName := "covid-19-project"
	collectionName := "new_cases_per_month_n_location"
	mt := mtest.New(t,
		mtest.NewOptions().ClientType(mtest.Mock),
		mtest.NewOptions().DatabaseName(databaseName),
		mtest.NewOptions().CollectionName(collectionName))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		newCase1 := services.NewCasesResponse{
			Year:       2021,
			Month:      12,
			Location:   "testLocation",
			TotalCases: 0,
			NewCases:   0,
		}
		newCase2 := services.NewCasesResponse{
			Year:       2022,
			Month:      1,
			Location:   "testLocation",
			TotalCases: 0,
			NewCases:   0,
		}
		expectedNewCases := []services.NewCasesResponse{newCase1, newCase2}

		log.Println("- Test - expectedNewCases: ", expectedNewCases[0])

		mockCursor1 := mtest.CreateCursorResponse(1, databaseName+"."+collectionName, mtest.FirstBatch, bson.D{
			primitive.E{Key: "_id", Value: primitive.NewObjectID()},
			primitive.E{Key: "year", Value: expectedNewCases[0].Year},
			primitive.E{Key: "month", Value: expectedNewCases[0].Month},
			primitive.E{Key: "location", Value: expectedNewCases[0].Location},
			primitive.E{Key: "total_cases", Value: expectedNewCases[0].TotalCases},
			primitive.E{Key: "new_cases", Value: expectedNewCases[0].NewCases},
		})
		mockCursor2 := mtest.CreateCursorResponse(1, databaseName+"."+collectionName, mtest.NextBatch, bson.D{
			primitive.E{Key: "_id", Value: primitive.NewObjectID()},
			primitive.E{Key: "year", Value: expectedNewCases[1].Year},
			primitive.E{Key: "month", Value: expectedNewCases[1].Month},
			primitive.E{Key: "location", Value: expectedNewCases[1].Location},
			primitive.E{Key: "total_cases", Value: expectedNewCases[1].TotalCases},
			primitive.E{Key: "new_cases", Value: expectedNewCases[1].NewCases},
		})
		killCursors := mtest.CreateCursorResponse(1, databaseName+"."+collectionName, mtest.NextBatch)

		mt.AddMockResponses(mockCursor1, mockCursor2, killCursors)
		repo := repo.NewRepository(mt.DB)
		actualNewCases, err := repo.GetNewCasesPerLocation("testLocation")

		assert.Equal(t, err, nil)
		assert.Equal(t, expectedNewCases, actualNewCases)
	})
}
