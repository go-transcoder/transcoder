package thumbnails

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

func TestExtractThumbs_CreateThumbs(t *testing.T) {
	var extractThumbs ExtractThumbs

	outputPath := fmt.Sprintf("%s/output", os.Getenv("LOCAL_STORAGE_PATH"))
	inputFilePath := fmt.Sprintf("%s/sample_test.mp4", os.Getenv("LOCAL_STORAGE_PATH"))

	err := extractThumbs.CreateThumbs(inputFilePath, outputPath)

	if err != nil {
		t.Fatalf("Error while extracting images err: %v", err)
	}
}
