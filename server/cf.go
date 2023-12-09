package main

import (
	"context"
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

func destination_address_list(noFilter bool, verified bool) ([]cloudflare.EmailRoutingDestinationAddress, error) {
	var destinations []cloudflare.EmailRoutingDestinationAddress
	api, err := initCFAPI()
	if err != nil {
		return nil, err
	}

	finished := false
	var result_info *cloudflare.ResultInfo

	for !finished {
		var params cloudflare.ListEmailRoutingAddressParameters
		if result_info == nil {
			if noFilter {
				params = cloudflare.ListEmailRoutingAddressParameters{}
			} else {
				params = cloudflare.ListEmailRoutingAddressParameters{Verified: &verified}
			}
		} else {
			if noFilter {
				params = cloudflare.ListEmailRoutingAddressParameters{ResultInfo: *result_info}
			} else {
				params = cloudflare.ListEmailRoutingAddressParameters{Verified: &verified, ResultInfo: *result_info}
			}
		}

		addresses, result_info, err := api.ListEmailRoutingDestinationAddresses(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), params)
		if err != nil {
			return nil, err
		}
		destinations = append(destinations, addresses...)
		if result_info.HasMorePages() {
			result_info.Next()
		} else {
			finished = true
		}
	}

	if err != nil {
		return nil, err
	}

	return destinations, nil
}
