package interfaces

import "github.com/go-transcoder/transcoder/internal/application/command"

type VideoService interface {
	Transcode(command *command.TranscodeCommand) (*command.TranscodeCommandResponse, error)
}
