package main

import (
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

// initCFAPI initializes the Cloudflare API client using the provided API key and email.
// It returns a pointer to the cloudflare.API struct and an error if any.
func initCFAPI() (*cloudflare.API, error) {
	api, err := cloudflare.New(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_API_EMAIL"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return api, nil
}
