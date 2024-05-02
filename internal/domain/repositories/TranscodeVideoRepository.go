package repositories

import (
	"github.com/go-transcoder/transcoder/internal/domain/entities"
	"github.com/google/uuid"
)

type TranscodeVideoRepository interface {
	Create(video *entities.TranscodeVideo) error
	FindById(id uuid.UUID) (*entities.TranscodeVideo, error)
	Update(product *entities.TranscodeVideo) error
}
