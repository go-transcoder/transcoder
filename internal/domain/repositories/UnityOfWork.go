package repositories

type UnityOfWork interface {
	StartTransaction()
	GetTranscodeVideosRepo() TranscodeVideoRepository
	Rollback()
	Commit()
}
