package server

import (
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/hillu/go-yara/v4"
)

// InitYARA initializes the YARA rules for the EMLServer.
// It compiles a YARA rule that detects malicious patterns in emails.
// The compiled rules are stored in the EMLServer's rules field.
// Returns an error if there was a problem compiling the rules.

func (s *EMLServer) InitYARA() error {
	yaraCompiler, err := yara.NewCompiler()
	if err != nil {
		return err
	}
	s.yc = yaraCompiler
	err = s.LoadYARADetections()
	if err != nil {
		return err
	}

	yaraRules, err := s.yc.GetRules()
	if err != nil {
		return err
	}
	s.rules = yaraRules
	return nil
}

// LoadYARADetections loads YARA detections from the R2 bucket.
// It downloads each detection file and adds its content to the YARA compiler.
func (s *EMLServer) LoadYARADetections() error {
	// Retrieve the R2 bucket name from an environment variable
	bucketName := os.Getenv("CLOUDFLARE_R2_YARA_BUCKET")

	// List objects in the R2 bucket
	listObjectsOutput, err := s.r2.ListObjectsV2(*s.r2ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return err
	}

	// Download each detection file into memory and add its content to the YARA compiler
	for _, detectionSummary := range listObjectsOutput.Contents {
		downloader := manager.NewDownloader(s.r2)
		downloadFile := manager.NewWriteAtBuffer([]byte{})

		_, err := downloader.Download(*s.r2ctx, downloadFile, &s3.GetObjectInput{
			Bucket: aws.String(os.Getenv("CLOUDFLARE_R2_YARA_BUCKET")),
			Key:    detectionSummary.Key,
		})
		if err != nil {
			return err
		}

		err = s.yc.AddString(string(downloadFile.Bytes()), "rules")
		if err != nil {
			return err
		}
	}

	return nil
}

// UploadYARADetection uploads a YARA detection to the R2 bucket.
// It takes the detection name and content as input and returns the upload information.
func (s *EMLServer) UploadYARADetection(detectionName string, detection string) (*manager.UploadOutput, error) {
	var uploadInfo *manager.UploadOutput

	// Create an uploader for the R2 client
	uploader := manager.NewUploader(s.r2)

	// Upload the detection file to the R2 bucket
	uploadInfo, err := uploader.Upload(*s.r2ctx, &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("CLOUDFLARE_R2_YARA_BUCKET")),
		Key:    aws.String(detectionName),
		Body:   strings.NewReader(detection),
	})
	if err != nil {
		return uploadInfo, err
	}

	return uploadInfo, nil
}
