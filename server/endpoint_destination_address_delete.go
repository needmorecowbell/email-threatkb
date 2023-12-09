package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// endpointDestinationAddressDeleteReq represents the request body for the endpointDestinationAddressDelete handler.
type endpointDestinationAddressDeleteReq struct {
	Email string `json:"email"`
}

// endpointDestinationAdressDeleteResp represents the response body for the endpointDestinationAddressDelete handler.
type endpointDestinationAdressDeleteResp GenericSuccessResponse

// endpointDestinationAddressDelete is the handler function for the DELETE /destination_address/delete endpoint.
// It deletes the specified destination address and returns the deleted destination address in the response.
func endpointDestinationAddressDelete(c *gin.Context) {
	var req endpointDestinationAddressDeleteReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenericErrorResponse{Message: "Bad request, must include email"})
		return
	}

	deleted_dest, err := destinationAddressDelete(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericErrorResponse{Message: fmt.Sprintf("Error deleting destination address: %s", err.Error())})
		return
	}
	c.JSON(http.StatusOK, endpointDestinationAdressDeleteResp{Message: fmt.Sprintf("Deleted %s as destination address", req.Email), Data: deleted_dest})
}
