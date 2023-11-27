// Package main provides the main entry point for the eml processor server.
// It includes functions for initializing the YARA compiler, scanning eml files for malicious patterns,
// and setting up the server router.
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hillu/go-yara/v4"
)

// initYARACompiler initializes the YARA compiler and adds a YARA rule for detecting malicious patterns.
// It returns the YARA compiler instance or an error if initialization fails.
func initYARACompiler() (*yara.Compiler, error) {
	yaraCompiler, err := yara.NewCompiler()
	if err != nil {
		return nil, err
	}
	err = yaraCompiler.AddString(`rule DetectMalicious {
		strings:
			$malicious_string = "malicious_pattern"
		condition:
			$malicious_string
	}`, "rules")

	if err != nil {
		return nil, err
	}
	return yaraCompiler, nil
}

// endpointScan is the handler function for the "/scan" endpoint.
// It scans the eml file received in the request body for malicious patterns using YARA rules.
// It returns the scan result as a JSON response.
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

// setupRouter sets up the server router with the necessary endpoints and handlers.
// It returns the configured gin.Engine instance.
func setupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Welcome to the eml processor"})
	})
	r.POST("/scan", endpointScan)
	return r
}

// main is the entry point for the eml processor server.
// It initializes the server router and starts the server.
func main() {
	log.Println("Starting eml processor")
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
