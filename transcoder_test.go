package main

import (
	"os"
	"testing"
)

func TestTranscoder_transcode(t *testing.T) {
	err := os.Setenv("FFMPEG_COMMAND_PATH", "/home/namir/Projects/my-transcoding/projects/transcoder/convert_video_cpu.sh")

	transcoder := transcoder{
		InputFile: "/home/namir/Projects/my-transcoding/projects/uploads/transcoder/sample_test.mp4",
		OutputDir: "/home/namir/Projects/my-transcoding/projects/uploads/transcoder/sample_test_cousa",
	}

	err = transcoder.transcode()

	if err != nil {
		t.Fatalf("unexpected error in transcode error:%v", err)
	}
}
