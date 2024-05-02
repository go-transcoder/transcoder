package dtos

import "github.com/aws/aws-sdk-go-v2/aws"

type BucketConfDto struct {
	S3Cfg    *aws.Config // Not sure about s3 configuration in the domain layer, Im thinking about abstracting the configuration
	S3Bucket string
}

func NewBucketConfDto(S3Cfg *aws.Config, S3Bucket string) *BucketConfDto {
	return &BucketConfDto{
		S3Cfg:    S3Cfg,
		S3Bucket: S3Bucket,
	}
}
