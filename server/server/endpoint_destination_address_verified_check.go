package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// endpointDestinationAddressVerifiedCheckReq represents the request structure for the EndpointDestinationAddressVerifiedCheck endpoint.
type endpointDestinationAddressVerifiedCheckReq struct {
	Email string `json:"email"`
}

// endpointDestinationAdressVerifiedCheckResp represents the response structure for the EndpointDestinationAddressVerifiedCheck endpoint.
type endpointDestinationAdressVerifiedCheckResp struct {
	GenericSuccessResponse
	Verified bool `json:"verified"`
}

// EndpointDestinationAddressVerifiedCheck is the handler function for the EndpointDestinationAddressVerifiedCheck endpoint.
// It checks if the destination address is verified and returns the result.
func (s *EMLServer) EndpointDestinationAddressVerifiedCheck(c *gin.Context) {
	var req endpointDestinationAddressVerifiedCheckReq
	var resp endpointDestinationAdressVerifiedCheckResp
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenericErrorResponse{Message: "Bad request, must include email"})
		return
	}

	verified, err := s.IsDestinationAddressVerified(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericErrorResponse{Message: fmt.Sprintf("Error deleting destination address: %s", err.Error())})
		return
	}
	resp.Verified = verified
	resp.Success = true
	resp.Message = "Destination address found"
	c.JSON(http.StatusOK, resp)
}
