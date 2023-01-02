package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/repositories"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/domain/values"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type filePartialPdfRepo struct {
	client     *s3.Client
	bucketName string
}

func NewFilePartialPdfRepo(c *s3.Client, b string) repositories.FilePartialPdfRepo {
	return &filePartialPdfRepo{client: c, bucketName: b}
}

func (r *filePartialPdfRepo) CreatePreSignedUrl(ctx context.Context, key values.FilePdfKey) (values.FilePreSignedUrl, error) {
	presignClient := s3.NewPresignClient(r.client)

	presignParams := &s3.GetObjectInput{
		Bucket: aws.String(r.bucketName),
		Key:    aws.String(string(key)), // TODO: Don't use string() to stringify value object. define valueObject.String() and use it.
	}

	// Apply an expiration via an option function
	presignDuration := func(po *s3.PresignOptions) {
		po.Expires = 60 * time.Minute
	}

	presignResult, err := presignClient.PresignGetObject(context.TODO(), presignParams, presignDuration)

	if err != nil {
		panic("Couldn't get presigned URL for GetObject")
	}

	fmt.Printf("Presigned URL For object: %s\n", presignResult.URL)
	return values.FilePreSignedUrl(presignResult.URL), nil
}
