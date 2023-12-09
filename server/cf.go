package main

import (
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

func initCFAPI() (*cloudflare.API, error) {
	api, err := cloudflare.New(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_API_EMAIL"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return api, nil
}
