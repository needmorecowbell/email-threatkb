package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

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

// endpointDestinationAddressList is the handler function for the /destination-addresses endpoint.
// It retrieves a list of destination addresses based on the provided filters.
func endpointDestinationAddressList(c *gin.Context) {
	var req endpointDestinationAddressListReq
	var resp endpointDestinationAdressListResp
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenericErrorResponse{Message: "Bad Request Error: Malformed request"})
		return
	}
	api, err := initCFAPI()
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericErrorResponse{Message: fmt.Sprintf("Error initializing Cloudflare API: %s", err.Error())})
		return
	}

	finished := false
	noFilter := req.NoFilter
	var result_info *cloudflare.ResultInfo

	for !finished {
		var params cloudflare.ListEmailRoutingAddressParameters
		if result_info == nil {
			if noFilter {
				params = cloudflare.ListEmailRoutingAddressParameters{}
			} else {
				params = cloudflare.ListEmailRoutingAddressParameters{Verified: &req.Verified}
			}
		} else {
			if noFilter {
				params = cloudflare.ListEmailRoutingAddressParameters{ResultInfo: *result_info}
			} else {
				params = cloudflare.ListEmailRoutingAddressParameters{Verified: &req.Verified, ResultInfo: *result_info}
			}
		}

		destinations, result_info, err := api.ListEmailRoutingDestinationAddresses(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, GenericErrorResponse{Message: fmt.Sprintf("Unable to list destinations: %s", err)})
			return
		}
		resp.Data = append(resp.Data, destinations...)
		if result_info.HasMorePages() {
			result_info.Next()
		} else {
			finished = true
		}
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericErrorResponse{Message: fmt.Sprintf("Error adding destination address: %s", err.Error())})
		return
	}

	if noFilter {
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
