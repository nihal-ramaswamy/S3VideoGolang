package dto

import (
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Data struct {
	Client  *s3.Client
	Context context.Context
	Bucket  string
}

func newClient(context context.Context, accessKey, secretAccessKey, region string) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(
		context,
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretAccessKey, "")),
	)
	if err != nil {
		return nil, err
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	return client, nil
}

func NewS3Data(bucket, accessKey, secretAccessKey, region string) (*S3Data, error) {
	context := context.TODO()
	client, err := newClient(context, accessKey, secretAccessKey, region)
	if nil != err {
		return nil, err
	}

	return &S3Data{
		Client:  client,
		Context: context,
		Bucket:  bucket,
	}, nil
}

func NewS3DataDefault(bucket, region, accessKey, secretAccessKey string) *S3Data {
	s3Data, err := NewS3Data(
		bucket,
		accessKey,
		secretAccessKey,
		region)
	if nil != err {
		panic(err)
	}

	return s3Data
}

func (s *S3Data) UploadFile(file multipart.File, filename string) error {
	_, err := s.Client.PutObject(s.Context, &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(filename),
		Body:   file,
	})

	return err
}
