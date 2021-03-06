package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		response := map[string]any{
			"status": 200,
			"data":   "API running smoothly",
		}
		c.JSON(http.StatusOK, response)
	}
}

func (server *Server) GetNewCasesPerLocation() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("- Handler - server is being executed")
		c.Header("Content-Type", "application/json")

		location := "Italy"
		data, err := server.service.GetNewCasesPerLocation(location)
		if err != nil {
			log.Println(err)
			return
		}

		response := map[string]any{
			"status": 200,
			"data":   data,
		}
		c.JSON(http.StatusOK, response)
	}

}
