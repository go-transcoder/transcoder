package main

import (
	"github.com/gofor-little/env"
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

	transcoder := transcoder{
		InputFile: "./resources/test/sample_test.mp4",
		OutputDir: "./resources/test/output/transcoded",
	}

	err := transcoder.transcode()

	if err != nil {
		t.Fatalf("unexpected error in transcode error:%v", err)
	}
}
