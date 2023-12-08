package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/gin-gonic/gin"
)

type endpointDestinationAddressAddReq struct {
	Email string
}

func endpointDestinationAddressAdd(c *gin.Context) {
	var req endpointDestinationAddressAddReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body"})
		return
	}
	api, err := initCFAPI()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Error initializing Cloudflare API"})
		return
	}
	dest_email, err := api.CreateEmailRoutingDestinationAddress(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), cloudflare.CreateEmailRoutingAddressParameters{Email: req.Email})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": fmt.Sprintf("Error adding destination address: %s", err.Error())})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": fmt.Sprintf("Added  destination address %s, verify address with email", dest_email.Email)})
}
