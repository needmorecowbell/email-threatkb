// Package main provides the main entry point for the eml processor server.
// It includes functions for initializing the YARA compiler, scanning eml files for malicious patterns,
// and setting up the server router.
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// setupRouter sets up the server router with the necessary endpoints and handlers.
// It returns the configured gin.Engine instance.
func setupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Welcome to the eml processor"})
	})
	r.POST("/destination_address/add", endpointDestinationAddressAdd)
	r.DELETE("/destination_address/delete", endpointDestinationAddressDelete)
	r.POST("/destination_address/get", endpointDestinationAddressGet)
	r.POST("/destination_address/list", endpointDestinationAddressList)
	r.POST("/scan", endpointScan)
	return r
}

// main is the entry point for the eml processor server.
// It initializes the server router and starts the server.
func main() {
	log.Println("Starting eml processor")
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
