package smil

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

func TestSmil_createSmilFile(t *testing.T) {
	var f CreateSmilFile

	outputPath := fmt.Sprintf("%s/output", os.Getenv("LOCAL_STORAGE_PATH"))

	err := f.CreateSmilFile(outputPath)

	if err != nil {
		t.Fatalf("Error while creating smil file err: %v", err)
	}
}
