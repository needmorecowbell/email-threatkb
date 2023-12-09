package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type endpointDestinationAddressDeleteReq struct {
	Email string `json:"email"`
}

type endpointDestinationAdressDeleteResp GenericSuccessResponse

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
