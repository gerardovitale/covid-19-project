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
func (s *Server) Run() error {
	// run function that initializes the routes
	r := s.Routes()

	// run the server through the router
	err := r.Run()

	if err != nil {
		log.Printf("- Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
