package ffmpeg

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

type CommandExec func(cv, inputFile, outputDir string) error

func (ffmpegApi CommandExec) FfmpegCommandExec(cv, inputFile, outputDir string) error {
	cmd := exec.Command(
		"/bin/bash",
		os.Getenv("CONVERT_FFMPEG_SCRIPT_PATH"),
		inputFile,
		outputDir,
		cv,
	)
	// Create buffers to capture output
	var stdout, stderr bytes.Buffer

	// Set the output buffers for the command
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()

	// Check for errors
	if err != nil {
		fmt.Println("ffmpeg stderr:", stderr.String())
		return err
	}

	// Print standard output
	fmt.Println("ffmpeg stdout:", stdout.String())

	return nil
}
