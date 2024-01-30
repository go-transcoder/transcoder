package ffmpeg

import (
	"os"
	"testing"
)

func TestFfmpegApi_FfmpegCommandExec(t *testing.T) {
	var ffmpegApi CommandExec

	err := ffmpegApi.FfmpegCommandExec("libx264", "/home/namir/Projects/my-transcoding/projects/uploads/transcoder/sample_test.mp4", "/home/namir/Projects/my-transcoding/projects/uploads/transcoder/sample_test_cousa")

	if err != nil {
		t.Fatalf("ffmpegCommandExec error: , %v", err)
	}

	// the output file should exist
	outputFile := "/home/namir/Projects/my-transcoding/projects/uploads/transcoder/sample_test_cousa"

	_, err = os.Stat(outputFile)

	if os.IsNotExist(err) {
		t.Fatalf("File %s is expected to be in path. error: %v", outputFile, err)
	}
}
