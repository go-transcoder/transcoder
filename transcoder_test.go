package main

import (
	"fmt"
	"github.com/gofor-little/env"
	"os"
	"testing"
)

func init() {
	// Set environment variables here
	// Load an .env file and set the key-value pairs as environment variables.
	if err := env.Load("./.env.test"); err != nil {
		panic(err)
	}
}

func TestTranscoder_transcode(t *testing.T) {

	outputPath := fmt.Sprintf("%s/output/transcoded", os.Getenv("LOCAL_STORAGE_PATH"))
	inputFilePath := fmt.Sprintf("%s/sample_test.mp4", os.Getenv("LOCAL_STORAGE_PATH"))

	transcoder := transcoder{
		InputFile: inputFilePath,
		OutputDir: outputPath,
	}

	err := transcoder.transcode()

	if err != nil {
		t.Fatalf("unexpected error in transcode error:%v", err)
	}
}
