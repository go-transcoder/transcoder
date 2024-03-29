package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofor-little/env"
	"os"
	"testing"
)

func init() {
	// Set environment variables here
	// Load an .env file and set the key-value pairs as environment variables.
	if err := env.Load("../../.env.test"); err != nil {
		panic(err)
	}
}

func TestS3BucketApi_UploadVideoDir(t *testing.T) {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		t.Fatalf("Error while loading config %v", err)
	}

	s3BucketApi := S3BucketApi{
		S3Client: s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = true
		}),
		BucketName: os.Getenv("INPUT_S3_BUCKET"),
	}

	s3BucketApi.UploadVideoDir("../../resources/test", "transcoded")
}

func TestS3BucketApi_DownloadFile(t *testing.T) {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		t.Fatalf("Error while loading config %v", err)
	}

	s3BucketApi := S3BucketApi{
		S3Client: s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = true
		}),
		BucketName: os.Getenv("INPUT_S3_BUCKET"),
	}

	downloadedFilePath := fmt.Sprintf("../../resources/test/%s", os.Getenv("OBJECT_NAME"))

	err = s3BucketApi.DownloadFile("uploads/sample-test.mp4", downloadedFilePath)

	_, err = os.Stat(downloadedFilePath)

	if os.IsNotExist(err) {
		t.Fatalf("File %s is expected to be in path. error: %v", downloadedFilePath, err)
	}
}
