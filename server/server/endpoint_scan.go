package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hillu/go-yara/v4"
)

func endpointScan(c *gin.Context) {
	// get the eml from the body
	eml_bytes, err := c.GetRawData()

	if err != nil {
		c.String(http.StatusBadRequest, "Error gathering eml from body")
		return
	}
	yc, err := initYARACompiler()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error compiling rules")
		return
	}
	yaraRules, err := yc.GetRules()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error retrieving compiled rules")
		return
	}

	scanner, err := yara.NewScanner(yaraRules)
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
	c.JSON(http.StatusOK, gin.H{"status": "malicious", "matches": matchStrings})

}
