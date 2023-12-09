package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// endpointMappingListResp represents the response structure for the EndpointMappingList API.
type endpointMappingListResp struct {
	GenericSuccessResponse
	Data []EmailMapping `json:"data"`
}

// EndpointMappingList is the handler function for the EndpointMappingList API.
// It retrieves the list of endpoint mappings and returns them as a JSON response.
func (s *EMLServer) EndpointMappingList(c *gin.Context) {
	var resp endpointMappingListResp

	mappings, err := s.MappingList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericErrorResponse{Message: fmt.Sprintf("Error adding destination address: %s", err.Error())})
		return
	}
	resp.Message = "Mappings retrieved"
	resp.Data = mappings
	c.JSON(http.StatusOK, resp)
}
