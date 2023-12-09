package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hillu/go-yara/v4"
)

// endpointScan is the handler function for the "/scan" endpoint.
// It scans the eml file received in the request body for malicious patterns using YARA rules.
// It returns the scan result as a JSON response.
func (s *EMLServer) EndpointScan(c *gin.Context) {
	// get the eml from the body
	eml_bytes, err := c.GetRawData()

	if err != nil {
		c.String(http.StatusBadRequest, "Error gathering eml from body")
		return
	}

	scanner, err := yara.NewScanner(s.rules)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating scanner from rules")
		return
	}

	var matches yara.MatchRules
	err = scanner.SetCallback(&matches).ScanMem(eml_bytes)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error scanning eml with rules")
		return
	}

	if len(matches) == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "clean", "matches": []string{}})
		return
	}

	var matchStrings []string
	for _, match := range matches {
		matchStrings = append(matchStrings, match.Rule)
	}

	// spin off a goroutine to send the eml to the cloudflare r2 bucket
	//go sendEMLToR2(eml_bytes)
	c.JSON(http.StatusOK, gin.H{"status": "malicious", "matches": matchStrings})

}
