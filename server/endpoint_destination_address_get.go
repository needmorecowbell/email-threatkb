package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// endpointDestinationAddressGetReq represents the request structure for the
// endpointDestinationAddressGet function.
type endpointDestinationAddressGetReq struct {
	Email string `json:"email"`
}

// endpointDestinationAdressGetResp represents the response structure for the
// endpointDestinationAddressGet function.
type endpointDestinationAdressGetResp GenericSuccessResponse

// endpointDestinationAddressGet is the handler function for the POST request to
// retrieve a destination address associated with a given email.
func endpointDestinationAddressGet(c *gin.Context) {
	var req endpointDestinationAddressGetReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenericErrorResponse{Message: "Bad request, must include email"})
		return
	}

	destinations, err := destinationAddressList(true, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericErrorResponse{Message: fmt.Sprintf("Error retrieving destination addresses: %s", err.Error())})
		return
	}
	for _, dest := range destinations {
		if dest.Email == req.Email {
			c.JSON(http.StatusOK, endpointDestinationAdressGetResp{Message: fmt.Sprintf("Retrieved destination address associated with %s", req.Email), Data: dest})
			return
		}
	}

	c.JSON(http.StatusNotFound, GenericErrorResponse{Message: fmt.Sprintf("Destination address %s not found", req.Email)})
}
