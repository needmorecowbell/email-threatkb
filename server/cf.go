package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

func initCFAPI() (*cloudflare.API, error) {
	// Construct a new API object using a global API key
	api, err := cloudflare.New(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_API_EMAIL"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Most API calls require a Context
	ctx := context.Background()

	// Fetch user details on the account
	u, err := api.UserDetails(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Print user details
	fmt.Println(u)
	return api, nil
}

// func sendEMLToR2(eml []byte) error {
// 	cf, err := initCFAPI()
// 	if err != nil {
// 		return err
// 	}
// 	ctx := context.Background()
// 	bucket,err:= cf.GetR2Bucket(ctx, cloudflare.AccountIdentifier(""), "emailvault")
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(bucket)
// 	bucket.

// }
