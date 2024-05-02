package entities

import (
	"errors"
	"fmt"
	"github.com/go-transcoder/transcoder/internal/domain/dtos"
	"github.com/go-transcoder/transcoder/internal/domain/events"
	videos_lib "github.com/go-transcoder/videos-lib"
	"github.com/google/uuid"
	"os"
	"strings"
	"time"
)

type TranscodeVideo struct {
	ID           uuid.UUID
	Title        string
	Path         string
	IsDownloaded bool
	IsTranscoded bool
	IsUploaded   bool
	Exception    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewTranscodeVideo(title string, path string) (*TranscodeVideo, error) {
	if title == "" || path == "" {
		return nil, errors.New("invalid details")
	}

	return &TranscodeVideo{
		ID:           uuid.New(),
		Title:        title,
		Path:         path,
		Exception:    "",
		IsDownloaded: false,
		IsTranscoded: false,
		IsUploaded:   false,
	}, nil
}

func (v *TranscodeVideo) AddException(err error) {
	v.Exception = fmt.Sprintf("Error: %v", err)
}

func (v *TranscodeVideo) SetIsUploaded() {
	v.IsUploaded = true
}

func (v *TranscodeVideo) SetIsDownloaded() {
	v.IsDownloaded = true
}

func (v *TranscodeVideo) SetIsTranscoded() {
	v.IsTranscoded = true
}

func (v *TranscodeVideo) GetDownloadLocalPath() string {
	storagePath := os.Getenv("STORAGE_PATH")

	if storagePath == "" {
		storagePath = "/tmp"
	}

	return fmt.Sprintf("%s/%s", storagePath, v.Title)
}

func (v *TranscodeVideo) GetTranscodeLocalPath() string {
	parts := strings.Split(v.Title, ".")

	storagePath := os.Getenv("STORAGE_PATH")

	if storagePath == "" {
		storagePath = "/tmp"
	}

	outputDir := fmt.Sprintf("%s/%s", storagePath, parts[0])

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, 0755)
	}

	return outputDir
}

func (v *TranscodeVideo) GetS3StoragePath() string {
	parts := strings.Split(v.Title, ".")

	return fmt.Sprintf("videos/%s", parts[0])
}

func (v *TranscodeVideo) Transcode(dto *dtos.BucketConfDto, eventService events.EventService) error {

	// Download
	err := videos_lib.DownloadFile(*dto.S3Cfg, dto.S3Bucket, v.Path, v.GetDownloadLocalPath())

	if err != nil {
		v.AddException(err)
		return err
	}

	v.SetIsDownloaded()

	// Transcode
	err = videos_lib.FfmpegTranscode(v.GetDownloadLocalPath(), v.GetTranscodeLocalPath())

	if err != nil {
		v.AddException(err)
		return err
	}

	err = videos_lib.CreateSmil(v.GetTranscodeLocalPath())

	if err != nil {
		v.AddException(err)
		return err
	}

	err = videos_lib.ExtractThumbnails(v.GetDownloadLocalPath(), v.GetTranscodeLocalPath())

	if err != nil {
		v.AddException(err)
		return err
	}

	v.SetIsTranscoded()

	// Upload
	err = videos_lib.UploadVideoDir(*dto.S3Cfg, dto.S3Bucket, v.GetTranscodeLocalPath(), v.GetS3StoragePath())

	if err != nil {
		v.AddException(err)
		return err
	}

	v.SetIsUploaded()

	// Delete original file
	err = videos_lib.DeleteFile(*dto.S3Cfg, dto.S3Bucket, v.Path)

	if err != nil {
		v.AddException(err)
		return err
	}

	// Dispatch the event
	KAFKATOPIC := os.Getenv("KAFKATOPIC")
	event := events.VideoHasBeenTranscoded{
		EventService: eventService,
		Topic:        KAFKATOPIC,
		VideoTitle:   v.Title,
	}
	err = event.Dispatch()
	if err != nil {
		return err
	}

	return nil
}
