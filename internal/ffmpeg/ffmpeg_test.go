package ffmpeg

import (
	"fmt"
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

func TestFfmpegApi_FfmpegCommandExec(t *testing.T) {
	var ffmpegApi CommandExec

	outputPath := fmt.Sprintf("%s/output", os.Getenv("LOCAL_STORAGE_PATH"))
	inputFilePath := fmt.Sprintf("%s/sample_test.mp4", os.Getenv("LOCAL_STORAGE_PATH"))

	err := ffmpegApi.FfmpegCommandExec(inputFilePath, outputPath)

	if err != nil {
		t.Fatalf("ffmpegCommandExec error: , %v", err)
	}

	// the output file should exist
	_, err = os.Stat(outputPath)

	if os.IsNotExist(err) {
		t.Fatalf("File %s is expected to be in path. error: %v", outputPath, err)
	}
}
