package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

type EmailMapping struct {
	GatewayAddress string `json:"gateway_address"`
	ForwardTo      string `json:"forward_to"`
	DateCreated    string `json:"date_created"`
}

func mappingList() ([]EmailMapping, error) {
	api, err := initCFAPI()
	if err != nil {
		return nil, err
	}
	keys, err := api.ListWorkersKVKeys(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), cloudflare.ListWorkersKVsParams{NamespaceID: os.Getenv("CLOUDFLARE_KV_NAMESPACE_ID")})
	if err != nil {
		return nil, err
	}
	var mappings []EmailMapping
	for _, key := range keys.Result {
		raw_mapping, err := api.GetWorkersKV(context.Background(), cloudflare.AccountIdentifier(os.Getenv("CLOUDFLARE_ACCOUNT_ID")), cloudflare.GetWorkersKVParams{NamespaceID: os.Getenv("CLOUDFLARE_KV_NAMESPACE_ID"), Key: key.Name})
		if err != nil {
			return nil, err
		}

		//marshal the raw mapping into a struct
		var mapping EmailMapping
		err = json.Unmarshal(raw_mapping, &mapping)
		if err != nil {
			return nil, err
		}

		mappings = append(mappings, mapping)
	}
	return mappings, nil
}
