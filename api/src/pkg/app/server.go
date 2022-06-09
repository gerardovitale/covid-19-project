package app

import (
	"covid-19-project/pkg/services"
	"log"

	"github.com/gin-gonic/gin"
)

// Server class definition
type Server struct {
	router  *gin.Engine
	service services.Service
}

// Server constructor
func NewServer(router *gin.Engine, service services.Service) *Server {
	return &Server{
		router:  router,
		service: service,
	}
}

// Server run main method
func (server *Server) Run() error {
	// run function that initializes the routes
	router := server.Routes()

	// run the server through the router
	err := router.Run()

	if err != nil {
		log.Printf("- Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}

// Server routes
func (server *Server) Routes() *gin.Engine {
	router := server.router

	// group all routes under /v1/api
	v1 := router.Group("/v1/api")
	{
		v1.GET("/status", server.ApiStatus())
		v1.GET("/new_cases", server.GetNewCasesPerLocation())
	}

	return router
}
