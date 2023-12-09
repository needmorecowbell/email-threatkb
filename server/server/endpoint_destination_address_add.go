package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// endpointDestinationAddressAddReq represents the request structure for adding a destination address.
type endpointDestinationAddressAddReq struct {
	Email string `json:"email"`
}

// endpointDestinationAdressAddResp represents the response structure for adding a destination address.
type endpointDestinationAdressAddResp GenericSuccessResponse

// EndpointDestinationAddressAdd is the handler function for the endpoint that adds a destination address.
// It receives a JSON request containing an email address and adds it as a destination address.
// If the request is invalid or there is an error adding the address, it returns an appropriate error response.
// Otherwise, it returns a success response with the added destination address.
func (s *EMLServer) EndpointDestinationAddressAdd(c *gin.Context) {
	var req endpointDestinationAddressAddReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenericErrorResponse{Message: "Bad request, must include email"})
		return
	}

	dest_email, err := s.DestinationAddressAdd(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericErrorResponse{Message: fmt.Sprintf("Error adding destination address: %s", err.Error())})
		return
	}
	c.JSON(http.StatusOK, endpointDestinationAdressAddResp{Message: fmt.Sprintf("Added %s as destination address, verify address with email", dest_email.Email), Data: dest_email})
}
