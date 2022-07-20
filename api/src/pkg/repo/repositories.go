package repo

import (
	"context"
	"log"

	"github.com/covid-19-project/api/src/pkg/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetNewCasesPerLocation(location string) ([]services.NewCasesResponse, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

// new_cases_per_month_n_location
func (repo *repository) GetNewCasesPerLocation(location string) ([]services.NewCasesResponse, error) {
	log.Println("- Repo - new cases repo is being executed")

	collection := repo.db.Collection("new_cases_per_month_n_location")
	filter := bson.D{primitive.E{Key: "location", Value: location}}
	cursor, err := collection.Find(context.TODO(), filter)

	log.Println("- Repo - mongoDb query was executed")

	if err != nil {
		panic(err)
	}

	var newCasesRecords []services.NewCasesResponse
	if err = cursor.All(context.TODO(), &newCasesRecords); err != nil {
		panic(err)
	}

	defer cursor.Close(context.TODO())

	return newCasesRecords, err

}
