package server

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// UploadEMLToVault uploads the given EML file to the vault.
// It takes the EML file key and the EML file content as input parameters.
// It returns the upload information and an error if any.
func (s *EMLServer) UploadEMLToVault(emlKey string, eml []byte) (*manager.UploadOutput, error) {
	fmt.Println("Uploading eml to vault: ", emlKey)
	var uploadInfo *manager.UploadOutput

	// Create an uploader for the R2 client
	uploader := manager.NewUploader(s.r2)

	// Upload the eml file to the R2 bucket
	uploadInfo, err := uploader.Upload(*s.r2ctx, &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("CLOUDFLARE_R2_VAULT_BUCKET")),
		Key:    aws.String(emlKey),
		Body:   bytes.NewReader(eml),
	})
	if err != nil {
		fmt.Printf("Error uploading eml to vault: %s", err)
		return uploadInfo, err
	}

	return uploadInfo, nil
}
