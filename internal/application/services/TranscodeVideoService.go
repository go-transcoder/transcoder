package services

import (
	"github.com/go-transcoder/transcoder/internal/application/command"
	"github.com/go-transcoder/transcoder/internal/domain/dtos"
	"github.com/go-transcoder/transcoder/internal/domain/entities"
	"github.com/go-transcoder/transcoder/internal/domain/events"
	"github.com/go-transcoder/transcoder/internal/domain/repositories"
)

type TranscodeVideoService struct {
	repo         repositories.TranscodeVideoRepository
	eventService events.EventService
}

func NewTranscodeVideoService(repo repositories.TranscodeVideoRepository, eventService events.EventService) *TranscodeVideoService {
	return &TranscodeVideoService{
		repo:         repo,
		eventService: eventService,
	}
}

func (service *TranscodeVideoService) Transcode(transcodeCommand *command.TranscodeCommand) (*command.TranscodeCommandResponse, error) {
	// Create the record in the database
	transcodeVideo, err := entities.NewTranscodeVideo(transcodeCommand.FileName, transcodeCommand.FilePath)

	if err != nil {
		return nil, err
	}

	transcodeErr := transcodeVideo.Transcode(dtos.NewBucketConfDto(transcodeCommand.S3Cfg, transcodeCommand.S3Bucket), service.eventService)

	// save the record in the database
	err = service.repo.Create(transcodeVideo)

	if transcodeErr != nil {
		return nil, transcodeErr
	}

	if err != nil {
		return nil, err
	}

	return &command.TranscodeCommandResponse{
		Message: "Successfully transcoded",
	}, nil

}
