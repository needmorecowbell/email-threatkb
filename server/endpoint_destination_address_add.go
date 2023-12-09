package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// endpointDestinationAddressAddReq represents the request payload for the endpointDestinationAddressAdd function.
type endpointDestinationAddressAddReq struct {
	Email string `json:"email"`
}

// endpointDestinationAdressAddResp represents the response payload for the endpointDestinationAddressAdd function.
type endpointDestinationAdressAddResp GenericSuccessResponse

// endpointDestinationAddressAdd is an HTTP endpoint that adds a destination address for email routing in Cloudflare.
// It expects a JSON payload with an "email" field specifying the email address to be added as a destination address.
// If the request is successful, it returns a JSON response with a success message and the created destination address.
// If there is an error during the process, it returns an appropriate error response.
func endpointDestinationAddressAdd(c *gin.Context) {
	var req endpointDestinationAddressAddReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenericErrorResponse{Message: "Bad request, must include email"})
		return
	}

	dest_email, err := destinationAddressAdd(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericErrorResponse{Message: fmt.Sprintf("Error adding destination address: %s", err.Error())})
		return
	}
	c.JSON(http.StatusOK, endpointDestinationAdressAddResp{Message: fmt.Sprintf("Added %s as destination address, verify address with email", dest_email.Email), Data: dest_email})
}
