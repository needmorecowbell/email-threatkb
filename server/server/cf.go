package server

import (
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

// InitCFAPI initializes the Cloudflare API client.
// It takes the Cloudflare API key and email from environment variables,
// creates a new Cloudflare API client, and assigns it to the EMLServer's cf field.
// If an error occurs during initialization, it returns the error.
func (s *EMLServer) InitCFAPI() error {
	api, err := cloudflare.New(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_API_EMAIL"))
	if err != nil {
		log.Fatal(err)
		return err
	}
	s.cf = api
	return nil
}
