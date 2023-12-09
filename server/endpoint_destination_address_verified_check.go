package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// endpointDestinationAddressDeleteReq represents the request body for the endpointDestinationAddressDelete handler.
type endpointDestinationAddressVerifiedCheckReq struct {
	Email string `json:"email"`
}

// endpointDestinationAdressDeleteResp represents the response body for the endpointDestinationAddressDelete handler.
type endpointDestinationAdressVerifiedCheckResp struct {
	GenericSuccessResponse
	Verified bool `json:"verified"`
}

// endpointDestinationAddressDelete is the handler function for the DELETE /destination_address/delete endpoint.
// It deletes the specified destination address and returns the deleted destination address in the response.
func endpointDestinationAddressVerifiedCheck(c *gin.Context) {
	var req endpointDestinationAddressVerifiedCheckReq
	var resp endpointDestinationAdressVerifiedCheckResp
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenericErrorResponse{Message: "Bad request, must include email"})
		return
	}

	verified, err := isDestinationAddressVerified(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericErrorResponse{Message: fmt.Sprintf("Error deleting destination address: %s", err.Error())})
		return
	}
	resp.Verified = verified
	resp.Success = true
	resp.Message = "Destination address found"
	c.JSON(http.StatusOK, resp)
}
