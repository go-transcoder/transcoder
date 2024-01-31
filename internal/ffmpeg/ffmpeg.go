package ffmpeg

import (
	"bytes"
	"fmt"
	"os/exec"
)

type CommandExec func(cv, inputFile, outputDir string) error

func (ffmpegApi CommandExec) FfmpegCommandExec(ffmpegCommandPath, inputFile, outputDir string) error {
	cmd := exec.Command(
		"/bin/bash",
		ffmpegCommandPath,
		inputFile,
		outputDir,
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
