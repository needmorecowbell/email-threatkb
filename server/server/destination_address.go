package server

import (
	"context"
	"errors"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

// DestinationAddressAdd adds a new destination address for email routing.
// It takes an email address as input and returns the created destination address and any error encountered.
func (s *EMLServer) DestinationAddressAdd(email string) (cloudflare.EmailRoutingDestinationAddress, error) {
	var dest_email cloudflare.EmailRoutingDestinationAddress

	dest_email, err := s.cf.CreateEmailRoutingDestinationAddress(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), cloudflare.CreateEmailRoutingAddressParameters{Email: email})
	if err != nil {
		return dest_email, err
	}
	return dest_email, nil
}

// DestinationAddressGet retrieves the destination address for the given email.
// It searches for the destination address in the list of destinations and returns it if found.
// If the destination address is not found, it returns an error.
func (s *EMLServer) DestinationAddressGet(email string) (cloudflare.EmailRoutingDestinationAddress, error) {
	var dest_email cloudflare.EmailRoutingDestinationAddress
	destinations, err := s.DestinationAddressList(true, false)
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

// IsDestinationAddressVerified checks if the destination address is verified.
// It retrieves the destination address for the given email and returns true if it is verified, false otherwise.
// An error is returned if there is any issue retrieving the destination address.
func (s *EMLServer) IsDestinationAddressVerified(email string) (bool, error) {
	destination, err := s.DestinationAddressGet(email)
	if err != nil {
		return false, err
	}
	return destination.Verified != nil, nil
}

// DestinationAddressDelete deletes the email routing destination address associated with the given email.
// It returns the deleted destination address and an error, if any.
func (s *EMLServer) DestinationAddressDelete(email string) (cloudflare.EmailRoutingDestinationAddress, error) {
	var dest_email cloudflare.EmailRoutingDestinationAddress
	destinations, err := s.DestinationAddressList(true, false)
	if err != nil {
		return dest_email, err
	}
	for _, dest := range destinations {
		if dest.Email == email {
			dest_email, err := s.cf.DeleteEmailRoutingDestinationAddress(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), dest.Tag)
			if err != nil {
				return dest_email, err
			}
			return dest_email, nil
		}
	}
	return dest_email, errors.New("destination address not found")
}

// DestinationAddressList returns a list of email routing destination addresses.
// If noFilter is true, it returns all destination addresses. If verified is true,
// it returns only the verified destination addresses.
func (s *EMLServer) DestinationAddressList(noFilter bool, verified bool) ([]cloudflare.EmailRoutingDestinationAddress, error) {
	var destinations []cloudflare.EmailRoutingDestinationAddress

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

		addresses, result_info, err := s.cf.ListEmailRoutingDestinationAddresses(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), params)
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

	return destinations, nil
}
