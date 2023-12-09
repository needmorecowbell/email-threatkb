// Package server provides the server setup for the eml processor.
package server

import (
	"context"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudflare/cloudflare-go"
	"github.com/gin-gonic/gin"
	"github.com/hillu/go-yara/v4"
)

type EMLServer struct {
	cf    *cloudflare.API
	yc    *yara.Compiler
	rules *yara.Rules
	r2    *s3.Client
	r2ctx *context.Context
}

// NewServer creates a new Server instance.
func NewServer() *EMLServer {
	var app = EMLServer{}
	err := app.InitCFAPI()
	if err != nil {
		panic(err)
	}

	err = app.InitR2()
	if err != nil {
		panic(err)
	}

	err = app.InitYARA()
	if err != nil {
		panic(err)
	}
	return &app
}

// SetupRouter sets up the server router with the necessary endpoints and handlers.
// It returns the configured gin.Engine instance.
func (s *EMLServer) SetupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Welcome to the eml processor"})
	})
	r.POST("/destination_address/add", s.EndpointDestinationAddressAdd)
	r.DELETE("/destination_address/delete", s.EndpointDestinationAddressDelete)
	r.POST("/destination_address/get", s.EndpointDestinationAddressGet)
	r.POST("/destination_address/list", s.EndpointDestinationAddressList)
	r.POST("/destination_address/verified", s.EndpointDestinationAddressVerifiedCheck)

	r.GET("/mapping/list", s.EndpointMappingList)

	r.POST("/scan", s.EndpointScan)
	r.POST("/yara/add", s.EndpointYARAAdd)

	return r
}
