package main

import (
	"context"
	"covid-19-project/pkg/app"
	"covid-19-project/pkg/repo"
	"covid-19-project/pkg/services"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	repository := setupRepositoryDB()
	service := services.NewService(repository)
	server := setupServer(service)

	err := server.Run()

	if err != nil {
		return err
	}

	return nil
}

func setupRepositoryDB() repo.Repository {
	uri := os.Getenv("MONGO_PASS")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	log.Println("- Main - Successfully connected and pinged to Mongo Atlas DB")

	return repo.NewRepository(client.Database("covid-19-project"))
} 

func setupServer(service services.Service) *app.Server {
	router := gin.Default()
	router.Use(cors.Default())

	return app.NewServer(router, service)
}
