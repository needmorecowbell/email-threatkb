package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// endpointDestinationAddressDeleteReq represents the request structure for deleting a destination address.
type endpointDestinationAddressDeleteReq struct {
	Email string `json:"email"`
}

// endpointDestinationAdressDeleteResp represents the response structure for deleting a destination address.
type endpointDestinationAdressDeleteResp GenericSuccessResponse

// EndpointDestinationAddressDelete is the handler function for deleting a destination address.
func (s *EMLServer) EndpointDestinationAddressDelete(c *gin.Context) {
	var req endpointDestinationAddressDeleteReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenericErrorResponse{Message: "Bad request, must include email"})
		return
	}

	deleted_dest, err := s.DestinationAddressDelete(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericErrorResponse{Message: fmt.Sprintf("Error deleting destination address: %s", err.Error())})
		return
	}
	c.JSON(http.StatusOK, endpointDestinationAdressDeleteResp{Message: fmt.Sprintf("Deleted %s as destination address", req.Email), Data: deleted_dest})
}
