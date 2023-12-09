package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// endpointDestinationAddressAddReq represents the request structure for adding a destination address.
type endpointYARAAddReq struct {
	Name string `json:"name"`
	Rule string `json:"rule"`
}

// endpointDestinationAdressAddResp represents the response structure for adding a destination address.
type endpointYARAAddResp GenericSuccessResponse

// EndpointDestinationAddressAdd is the handler function for the endpoint that adds a destination address.
// It receives a JSON request containing an email address and adds it as a destination address.
// If the request is invalid or there is an error adding the address, it returns an appropriate error response.
// Otherwise, it returns a success response with the added destination address.
func (s *EMLServer) EndpointYARAAdd(c *gin.Context) {
	var req endpointYARAAddReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenericErrorResponse{Message: "Bad request, must include email"})
		return
	}

	resp, err := s.UploadYARADetection(req.Name, req.Rule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericErrorResponse{Message: fmt.Sprintf("Error adding yara rule: %s", err.Error())})
		return
	}
	c.JSON(http.StatusOK, endpointYARAAddResp{Success: true, Message: fmt.Sprintf("Added %s to yara corpus, location: %s", *resp.Key, resp.Location), Data: resp})
}
