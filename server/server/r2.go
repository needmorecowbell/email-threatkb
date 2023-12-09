package server

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// InitR2 initializes the R2 client for Cloudflare R2 storage.
// It sets up the AWS SDK configuration and creates an S3 client.
// It also lists the objects in the R2 bucket and prints their details.
func (s *EMLServer) InitR2() error {
	// Retrieve Cloudflare account ID, R2 access key ID, and R2 secret access key from environment variables
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	accessKeyID := os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("CLOUDFLARE_R2_SECRET_ACCESS_KEY")

	// Define a custom endpoint resolver for R2 storage
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID),
			HostnameImmutable: true,
			Source:            aws.EndpointSourceCustom,
		}, nil
	})

	// Create a context
	ctx := context.TODO()
	s.r2ctx = &ctx

	// Load the default AWS SDK configuration with the custom endpoint resolver and static credentials provider
	cfg, err := config.LoadDefaultConfig(*s.r2ctx,
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Create an S3 client using the loaded configuration
	client := s3.NewFromConfig(cfg)
	s.r2 = client

	return nil
}
