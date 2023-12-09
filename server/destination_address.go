package main

import (
	"context"
	"errors"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

// destinationAddressAdd adds a new destination email address for email routing.
// It takes an email address as input and returns the created destination address and any error encountered.
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

// destinationAddressGet retrieves the email routing destination address for the given email.
// It searches for the destination address in the list of destinations and returns it if found.
// If the destination address is not found, it returns an error.
func destinationAddressGet(email string) (cloudflare.EmailRoutingDestinationAddress, error) {
	var dest_email cloudflare.EmailRoutingDestinationAddress
	destinations, err := destinationAddressList(true, false)
	if err != nil {
		return dest_email, err
	}
	for _, dest := range destinations {
		if dest.Email == email {
			return dest, nil
		}
	}
	return dest_email, errors.New("destination address not found")
}

// isDestinationAddressVerified checks if the destination address is verified.
// It retrieves the destination address using the provided email and returns
// true if the address is verified, false otherwise. An error is returned if
// there was a problem retrieving the destination address.
func isDestinationAddressVerified(email string) (bool, error) {
	destination, err := destinationAddressGet(email)
	if err != nil {
		return false, err
	}
	return destination.Verified != nil, nil
}

// destinationAddressDelete deletes the email routing destination address associated with the given email.
// It returns the deleted destination address and an error, if any.
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

// destinationAddressList retrieves a list of email routing destination addresses.
// It takes two parameters: noFilter (a boolean indicating whether to apply any filters) and verified (a boolean indicating whether to filter by verified addresses).
// It returns a slice of cloudflare.EmailRoutingDestinationAddress and an error.
// If the initialization of the Cloudflare API fails, it returns an error.
// It iterates through the paginated results until all addresses are retrieved.
// If any error occurs during the retrieval process, it returns an error.
// Otherwise, it returns the list of destination addresses.
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
