package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/gin-gonic/gin"
)

type endpointDestinationAddressDeleteReq struct {
	Email string `json:"email"`
}

type endpointDestinationAdressDeleteResp GenericSuccessResponse

func endpointDestinationAddressDelete(c *gin.Context) {
	var req endpointDestinationAddressDeleteReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenericErrorResponse{Message: "Bad request, must include email"})
		return
	}

	destinations, err := destination_address_list(true, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericErrorResponse{Message: fmt.Sprintf("Error retrieving destination addresses: %s", err.Error())})
		return
	}
	for _, dest := range destinations {
		if dest.Email == req.Email {
			api, err := initCFAPI()
			if err != nil {
				c.JSON(http.StatusInternalServerError, GenericErrorResponse{Message: fmt.Sprintf("Error initializing Cloudflare API: %s", err.Error())})
				return
			}
			deleted_item, err := api.DeleteEmailRoutingDestinationAddress(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), dest.Tag)
			if err != nil {
				c.JSON(http.StatusInternalServerError, GenericErrorResponse{Message: fmt.Sprintf("Error deleting destination address: %s", err.Error())})
				return
			}
			c.JSON(http.StatusOK, endpointDestinationAdressDeleteResp{Message: fmt.Sprintf("Deleted %s as destination address", req.Email), Data: deleted_item})
			return
		}
	}

	c.JSON(http.StatusNotFound, GenericErrorResponse{Message: fmt.Sprintf("Destination address %s not found", req.Email)})
}
