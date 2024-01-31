package smil

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

func TestSmil_createSmilFile(t *testing.T) {
	var f CreateSmilFile

	err := f.CreateSmilFile("../../resources/test/output")

	if err != nil {
		t.Fatalf("Error while creating smil file err: %v", err)
	}
}
