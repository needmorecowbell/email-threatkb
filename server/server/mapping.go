package server

import (
	"context"
	"encoding/json"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

// EmailMapping represents the mapping between a gateway address and the email address it forwards to.
type EmailMapping struct {
	GatewayAddress string `json:"gateway_address"` // The gateway address that receives the emails.
	ForwardTo      string `json:"forward_to"`      // The email address to which the emails are forwarded.
	DateCreated    string `json:"date_created"`    // The date when the mapping was created.
}

// MappingList returns a list of EmailMapping objects by retrieving the keys from the Cloudflare Workers KV store
// and unmarshaling the corresponding values into EmailMapping structs.
// It uses the Cloudflare API to interact with the KV store.
// The Cloudflare account ID and KV namespace ID are retrieved from environment variables.
// If an error occurs during the retrieval or unmarshaling process, it is returned along with a nil slice of EmailMapping.
func (s *EMLServer) MappingList() ([]EmailMapping, error) {

	keys, err := s.cf.ListWorkersKVKeys(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), cloudflare.ListWorkersKVsParams{NamespaceID: os.Getenv("CLOUDFLARE_KV_NAMESPACE_ID")})
	if err != nil {
		return nil, err
	}
	var mappings []EmailMapping
	for _, key := range keys.Result {
		raw_mapping, err := s.cf.GetWorkersKV(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), cloudflare.GetWorkersKVParams{NamespaceID: os.Getenv("CLOUDFLARE_KV_NAMESPACE_ID"), Key: key.Name})
		if err != nil {
			return nil, err
		}

		var mapping EmailMapping
		err = json.Unmarshal(raw_mapping, &mapping)
		if err != nil {
			return nil, err
		}

		mappings = append(mappings, mapping)
	}
	return mappings, nil
}
