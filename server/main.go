package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	r := gin.Default()

	// Ping test
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Welcome to the eml processor"})
	})
	r.POST("/scan", server.endpointScan)
	return r
}

func main() {
	log.Println("Starting eml processor")
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
