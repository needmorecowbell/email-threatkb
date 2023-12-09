package main

import (
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go"
	"github.com/gin-gonic/gin"
)

// endpointDestinationAddressListReq represents the request body for the endpointDestinationAddressList handler.
type endpointDestinationAddressListReq struct {
	Verified bool `json:"verified,omitempty"`
	NoFilter bool `json:"no_filter,omitempty"`
}

// endpointDestinationAdressListResp represents the response body for the endpointDestinationAddressList handler.
type endpointDestinationAdressListResp struct {
	GenericSuccessResponse
	Data []cloudflare.EmailRoutingDestinationAddress `json:"data"`
}

// endpointDestinationAddressList retrieves a list of destination addresses based on the provided filters.
func endpointDestinationAddressList(c *gin.Context) {
	var req endpointDestinationAddressListReq
	var resp endpointDestinationAdressListResp
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenericErrorResponse{Message: "Bad Request Error: Malformed request"})
		return
	}

	destinations, err := destinationAddressList(req.NoFilter, req.Verified)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericErrorResponse{Message: fmt.Sprintf("Error retrieving destination addresses: %s", err.Error())})
		return
	}
	resp.Data = destinations

	if req.NoFilter {
		resp.Message = "Retrieved all destination addresses"
	} else {
		if req.Verified {
			resp.Message = "Retrieved all verified destination addresses"
		} else {
			resp.Message = "Retrieved all unverified destination addresses"
		}
	}
	c.JSON(http.StatusOK, resp)
}
