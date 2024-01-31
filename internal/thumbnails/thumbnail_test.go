package thumbnails

import (
	"github.com/gofor-little/env"
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

	err := extractThumbs.CreateThumbs("../../resources/test/sample_test.mp4", "../../resources/test/output")

	if err != nil {
		t.Fatalf("Error while extracting images err: %v", err)
	}
}
