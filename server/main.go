// Package main provides the main entry point for the eml processor server.
// It includes functions for initializing the YARA compiler, scanning eml files for malicious patterns,
// and setting up the server router.
package main

import (
	"fmt"
	"os"

	"github.com/needmorecowbell/email_threatkb/server/server"
)

// main is the entry point for the eml processor server.
// It initializes the server router and starts the server.
func main() {
	s := server.NewServer()
	r := s.SetupRouter()
	r.Run(fmt.Sprintf(":%s", os.Getenv("EML_SERVER_PORT")))
}
