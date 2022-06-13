package app

import (
	"os"
	"testing"

	"github.com/covid-19-project/api/src/pkg/repo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupTestServer(t *testing.T, repo repo.Repository) *Server {
	router := gin.Default()
	router.Use(cors.Default())
	server := NewServer(router, repo)
	server.Routes()

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
