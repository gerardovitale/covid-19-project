package main

import (
	"covid-19-project/pkg/app"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	// create router dependency
	router := gin.Default()
	router.Use(cors.Default())

	server := app.NewServer(router)

	// start the server
	err := server.Run()

	if err != nil {
		return err
	}

	return nil
}
