package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/covid-19-project/api/src/pkg/app"
	"github.com/covid-19-project/api/src/pkg/repo"
	"github.com/covid-19-project/api/src/pkg/services"
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

	repository := repo.NewRepository(client.Database("covid-19-project"))
	service := services.NewService(repository)

	// create router dependency
	router := gin.Default()
	router.Use(cors.Default())

	server := app.NewServer(router, service)

	// start the server
	err = server.Run()

	if err != nil {
		return err
	}

	return nil
}
