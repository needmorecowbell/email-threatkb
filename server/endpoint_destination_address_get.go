package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type endpointDestinationAddressGetReq struct {
	Email string `json:"email"`
}

type endpointDestinationAdressGetResp GenericSuccessResponse

func endpointDestinationAddressGet(c *gin.Context) {
	var req endpointDestinationAddressGetReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenericErrorResponse{Message: "Bad request, must include email"})
		return
	}

	destinations, err := destinationAddressList(true, false)
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
