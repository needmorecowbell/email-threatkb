package main

import (
	"context"
	"errors"
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

func destinationAddressAdd(email string) (cloudflare.EmailRoutingDestinationAddress, error) {
	var dest_email cloudflare.EmailRoutingDestinationAddress

	api, err := initCFAPI()
	if err != nil {
		return dest_email, err
	}
	dest_email, err = api.CreateEmailRoutingDestinationAddress(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), cloudflare.CreateEmailRoutingAddressParameters{Email: email})
	if err != nil {
		return dest_email, err
	}
	return dest_email, nil
}

func destinationAddressDelete(email string) (cloudflare.EmailRoutingDestinationAddress, error) {
	var dest_email cloudflare.EmailRoutingDestinationAddress
	destinations, err := destinationAddressList(true, false)
	if err != nil {
		return dest_email, err
	}
	for _, dest := range destinations {
		if dest.Email == email {
			api, err := initCFAPI()
			if err != nil {
				return dest_email, err
			}
			dest_email, err := api.DeleteEmailRoutingDestinationAddress(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), dest.Tag)
			if err != nil {
				return dest_email, err
			}
			return dest_email, nil
		}
	}
	return dest_email, errors.New("destination address not found")

}

func destinationAddressList(noFilter bool, verified bool) ([]cloudflare.EmailRoutingDestinationAddress, error) {
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
