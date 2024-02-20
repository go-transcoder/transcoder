package s3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"log"
	"os"
	"path/filepath"
)

type S3BucketApi struct {
	S3Client   *s3.Client
	BucketName string
}

func (s3BucketApi S3BucketApi) DownloadFile(objectKey string, fileName string) error {
	result, err := s3BucketApi.S3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s3BucketApi.BucketName),
		Key:    aws.String(objectKey),
	})

	if err != nil {
		log.Printf("Couldn't get object %v:%v. Here's why: %v\n", s3BucketApi.BucketName, objectKey, err)
		return err
	}

	file, err := os.Create(fileName)

	if err != nil {
		log.Printf("Couldn't create file %v. Here's why: %v\n", fileName, err)
		return err
	}

	defer file.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Printf("Couldn't read object body from %v. Here's why: %v\n", objectKey, err)
	}
	_, err = file.Write(body)
	return err
}

func (s3BucketApi S3BucketApi) UploadVideoDir(videoDirPath string, prefix string) error {
	walker := make(fileWalk)
	go func() {
		// Gather the files to videos by walking the path recursively
		if err := filepath.Walk(videoDirPath, walker.Walk); err != nil {
			log.Fatalln("Walk failed:", err)
		}
		close(walker)
	}()

	for path := range walker {
		rel, err := filepath.Rel(videoDirPath, path)

		if err != nil {
			log.Fatalln("Unable to get relative path:", path, err)
		}
		file, err := os.Open(path)

		if err != nil {
			log.Println("Failed opening file", path, err)
			continue
		}

		defer file.Close()

		result, err := s3BucketApi.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(s3BucketApi.BucketName),
			Key:    aws.String(filepath.Join(prefix, rel)),
			Body:   file,
		})
		if err != nil {
			log.Fatalln("Failed to videos", path, err)
		}
		log.Println("Uploaded", path, result.ResultMetadata)
	}
	return nil
}
