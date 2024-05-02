package command

import "github.com/aws/aws-sdk-go-v2/aws"

type TranscodeCommand struct {
	FileName string
	FilePath string
	S3Cfg    *aws.Config
	S3Bucket string
}
