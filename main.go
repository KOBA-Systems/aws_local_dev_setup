package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// CustomEvent for lambda
type CustomEvent struct {
	ID   string
	Name string
}

var (
	awsRegion   string
	awsEndpoint string
	bucketName  string

	s3svc *s3.Client
)

func init() {
	awsRegion = os.Getenv("AWS_REGION")
	awsEndpoint = os.Getenv("AWS_ENDPOINT")
	bucketName = os.Getenv("S3_BUCKET")

	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: awsRegion,
			}, nil
		}

		// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithEndpointResolver(customResolver),
	)
	if err != nil {
		log.Fatalf("Cannot load the AWS configs: %s", err)
	}

	s3svc = s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
}

func handler(ctx context.Context, event CustomEvent) error {
	s3Key := fmt.Sprintf("%s.txt", event.ID)
	body := []byte(fmt.Sprintf("Hello, %s", event.Name))
	resp, err := s3svc.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:             aws.String(bucketName),
		Key:                aws.String(s3Key),
		Body:               bytes.NewReader(body),
		ContentLength:      int64(len(body)),
		ContentType:        aws.String("application/text"),
		ContentDisposition: aws.String("attachment"),
	})
	log.Printf("S3 PutObject response: %+v", resp)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
