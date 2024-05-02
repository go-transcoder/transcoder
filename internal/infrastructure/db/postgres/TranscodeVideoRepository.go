package postgres

import (
	"github.com/go-transcoder/transcoder/internal/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TranscodeVideoRepository struct {
	db *gorm.DB
}

func NewTranscodeVideoRepo(db *gorm.DB) *TranscodeVideoRepository {
	return &TranscodeVideoRepository{
		db: db,
	}
}

func (repo *TranscodeVideoRepository) Create(video *entities.TranscodeVideo) error {

	if err := repo.db.Create(ToDBTranscodeVideo(video)).Error; err != nil {
		return err
	}

	storedVideo, err := repo.FindById(video.ID)

	if err != nil {
		return err
	}

	*video = *storedVideo

	return nil

}

func (repo *TranscodeVideoRepository) FindById(id uuid.UUID) (*entities.TranscodeVideo, error) {
	var transcodeVideo TranscodeVideo

	if err := repo.db.Find(&transcodeVideo, id).Error; err != nil {
		return nil, err
	}

	return FromDBTranscodeVideo(&transcodeVideo)
}

func (repo *TranscodeVideoRepository) Update(video *entities.TranscodeVideo) error {
	dbRecord := ToDBTranscodeVideo(video)

	err := repo.db.Model(TranscodeVideo{}).Where("id = ?", dbRecord.ID).Updates(dbRecord).Error
	if err != nil {
		return err
	}

	// Read row from DB to never return different data than persisted
	storedVideo, err := repo.FindById(dbRecord.ID)
	if err != nil {
		return err
	}

	// Map back to domain entity
	*video = *storedVideo

	return nil

}
