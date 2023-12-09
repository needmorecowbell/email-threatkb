package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// endpointDestinationAddressGetReq represents the request structure for the EndpointDestinationAddressGet handler.
type endpointDestinationAddressGetReq struct {
	Email string `json:"email"`
}

// endpointDestinationAdressGetResp represents the response structure for the EndpointDestinationAddressGet handler.
type endpointDestinationAdressGetResp GenericSuccessResponse

// EndpointDestinationAddressGet is the handler function for the GET request to retrieve a destination address.
// It retrieves the destination addresses and checks if the requested email is associated with any of them.
// If found, it returns the associated destination address.
// If not found, it returns a not found error.
func (s *EMLServer) EndpointDestinationAddressGet(c *gin.Context) {
	var req endpointDestinationAddressGetReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenericErrorResponse{Message: "Bad request, must include email"})
		return
	}

	destinations, err := s.DestinationAddressList(true, false)
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
