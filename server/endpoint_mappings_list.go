package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type endpointMappingListResp struct {
	GenericSuccessResponse
	Data []EmailMapping `json:"data"`
}

func endpointMappingList(c *gin.Context) {
	var resp endpointMappingListResp

	mappings, err := mappingList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenericErrorResponse{Message: fmt.Sprintf("Error adding destination address: %s", err.Error())})
		return
	}
	resp.Message = "Mappings retrieved"
	resp.Data = mappings
	c.JSON(http.StatusOK, resp)
}
