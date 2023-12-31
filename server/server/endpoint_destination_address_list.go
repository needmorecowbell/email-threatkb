package server

import (
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go"
	"github.com/gin-gonic/gin"
)

// endpointDestinationAddressListReq represents the request body for the EndpointDestinationAddressList endpoint.
type endpointDestinationAddressListReq struct {
	Verified bool `json:"verified,omitempty"`
	NoFilter bool `json:"no_filter,omitempty"`
}

// endpointDestinationAdressListResp represents the response body for the EndpointDestinationAddressList endpoint.
type endpointDestinationAdressListResp struct {
	GenericSuccessResponse
	Data []cloudflare.EmailRoutingDestinationAddress `json:"data"`
}

// EndpointDestinationAddressList is the handler function for the EndpointDestinationAddressList endpoint.
// It retrieves the list of destination addresses based on the provided request parameters.
func (s *EMLServer) EndpointDestinationAddressList(c *gin.Context) {
	var req endpointDestinationAddressListReq
	var resp endpointDestinationAdressListResp

	// Bind the request body to the req struct
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenericErrorResponse{Message: "Bad Request Error: Malformed request"})
		return
	}

	// Retrieve the destination addresses based on the request parameters
	destinations, err := s.DestinationAddressList(req.NoFilter, req.Verified)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericErrorResponse{Message: fmt.Sprintf("Error retrieving destination addresses: %s", err.Error())})
		return
	}
	resp.Data = destinations

	// Set the response message based on the request parameters
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
