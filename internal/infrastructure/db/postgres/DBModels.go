package postgres

import (
	"github.com/google/uuid"
	"time"
)

type TranscodeVideo struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	Title        string
	IsDownloaded bool
	IsTranscoded bool
	IsUploaded   bool
	Exception    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
