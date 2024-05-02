package postgres

import "github.com/go-transcoder/transcoder/internal/domain/entities"

func ToDBTranscodeVideo(video *entities.TranscodeVideo) *TranscodeVideo {
	var v = &TranscodeVideo{
		Title:        video.Title,
		IsDownloaded: video.IsDownloaded,
		IsTranscoded: video.IsTranscoded,
		IsUploaded:   video.IsUploaded,
		Exception:    video.Exception,
		CreatedAt:    video.CreatedAt,
		UpdatedAt:    video.UpdatedAt,
	}

	v.ID = video.ID

	return v
}

func FromDBTranscodeVideo(dbProduct *TranscodeVideo) (*entities.TranscodeVideo, error) {
	var v = &entities.TranscodeVideo{
		Title:        dbProduct.Title,
		IsDownloaded: dbProduct.IsDownloaded,
		IsTranscoded: dbProduct.IsTranscoded,
		IsUploaded:   dbProduct.IsUploaded,
		Exception:    dbProduct.Exception,
		CreatedAt:    dbProduct.CreatedAt,
		UpdatedAt:    dbProduct.UpdatedAt,
	}
	v.ID = dbProduct.ID

	return v, nil
}
