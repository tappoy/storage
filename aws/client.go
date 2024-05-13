package aws

import (
	"github.com/tappoy/archive"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	client *s3.Client
}

// NewClient is a factory method for S3Client
//
// Errors:
func NewClient(region, accessKey, secretKey, endpoint string) (archive.Client, error) {
	appCreds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""))

	cfg := aws.Config{
		Region:       region,
		Credentials:  appCreds,
		BaseEndpoint: &endpoint,
	}

	return S3Client{client: s3.NewFromConfig(cfg)}, nil
}
