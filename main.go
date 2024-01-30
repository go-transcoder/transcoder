package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3_internal "github.com/go-transcoder/transcoder/internal/s3"
	"log"
	"os"
)

func main() {
	fmt.Println("The transcoder in action")

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	s3BucketApi := s3_internal.S3BucketApi{
		S3Client:   s3.NewFromConfig(cfg),
		BucketName: os.Getenv("INPUT_S3_BUCKET"),
	}

	objectPath := os.Getenv("UPLOADER_APP_UPLOAD_PATH") + "/" + os.Getenv("OBJECT_NAME")

	// download the video
	// path of the video that it will be downloaded to
	err = s3BucketApi.DownloadFile(fmt.Sprintf("uploads/%s", os.Getenv("OBJECT_NAME")), objectPath)

	if err != nil {
		fmt.Printf("Error while Downloading err: %v", err)
	}

	transcoderApi := transcoder{
		InputFile: objectPath,
		OutputDir: fmt.Sprintf("/tmp/%s", os.Getenv("OBJECT_NAME")),
	}

	// transcode the video
	err = transcoderApi.transcode()

	if err != nil {
		fmt.Printf("Error while transcoding err: %v", err)
	}
	// upload back to s3
}
