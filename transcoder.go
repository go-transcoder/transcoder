package main

import (
	"fmt"
	"github.com/go-transcoder/transcoder/internal/ffmpeg"
	"github.com/go-transcoder/transcoder/internal/smil"
	"github.com/go-transcoder/transcoder/internal/thumbnails"
	"os"
	"sync"
)

type transcoder struct {
	InputFile string
	OutputDir string
}

func (transcoder transcoder) transcode() error {
	// make sure that the directory exists
	_, err := os.Stat(transcoder.OutputDir)

	if os.IsNotExist(err) {
		err := os.Mkdir(transcoder.OutputDir, 0755)
		if err != nil {
			return err
		}
	}

	transcodingChannel := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		var ffmpegApi ffmpeg.CommandExec
		err := ffmpegApi.FfmpegCommandExec("libx264", transcoder.InputFile, transcoder.OutputDir)

		if err != nil {
			fmt.Printf("error while executing ffmpeg command: %v", err)
			transcodingChannel <- "error while executing ffmpeg command"
		} else {
			transcodingChannel <- fmt.Sprintf("ffmpeg finished transcoding: %s", transcoder.InputFile)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		var smilApi smil.CreateSmilFile
		err := smilApi.CreateSmilFile(transcoder.OutputDir)

		if err != nil {
			fmt.Printf("error while generating smil file: %v", err)
			transcodingChannel <- "error while generating smil file"
		} else {
			transcodingChannel <- fmt.Sprintf("finished generating smil")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		var thumbnailsApi thumbnails.ExtractThumbs
		err := thumbnailsApi.CreateThumbs(transcoder.InputFile, transcoder.OutputDir)

		if err != nil {
			fmt.Printf("error while extracting images: %v", err)
			transcodingChannel <- "error while extracting images"
		} else {
			transcodingChannel <- fmt.Sprintf("images are extracted")
		}

	}()

	go func() {
		for i := range transcodingChannel {
			fmt.Println(i)
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	close(transcodingChannel)
	return nil
}
