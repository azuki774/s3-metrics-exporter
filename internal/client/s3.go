package client

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/s3-metrics-exporter/internal/model"
)

type S3Client struct {
	Svc  *s3.S3
	Sess *session.Session
}

func NewS3Client() (*S3Client, error) {
	endpoint := os.Getenv("ENDPOINT")     // ex. https://s3.ap-northeast-1.wasabisys.com or empty(default S3)
	region := os.Getenv("ap-northeast-1") // ex. ap-northeast-1
	creds := credentials.NewSharedCredentials("", "default")
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String(region),
		Endpoint:    &endpoint,
	})
	if err != nil {
		return nil, err
	}

	Svc := s3.New(sess)
	return &S3Client{Svc: Svc, Sess: sess}, nil
}

func (s *S3Client) GetFileInfo(ctx context.Context, objInfo *model.ObjectInfo) (err error) {
	obj, err := s.Svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(objInfo.BucketName),
		Key:    aws.String(objInfo.Key),
	})

	if err != nil {
		return err
	}

	if obj.LastModified != nil {
		objInfo.LastModified = *obj.LastModified
	}
	if obj.ContentLength != nil {
		objInfo.ContentLength = *obj.ContentLength
	}

	objInfo.LastUpdated = time.Now()
	return nil
}