package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Config struct {
	ID               string
	Secret           string
	Token            string
	Region           string
	Origin           string
	S3ForcePathStyle bool
}

func NewSession(cnf S3Config) (*s3.Client, error) {
	cred := credentials.NewStaticCredentialsProvider(cnf.ID, cnf.Secret, cnf.Token)
	endpointResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{URL: cnf.Origin}, nil
	})
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(cnf.Region),
		config.WithCredentialsProvider(cred),
		config.WithEndpointResolverWithOptions(endpointResolver),
	)
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.UsePathStyle = cnf.S3ForcePathStyle
	})
	return client, nil
}
