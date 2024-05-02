package events

import (
	"encoding/json"
	"os"
)

type VideoHasBeenTranscoded struct {
	VideoTitle   string
	EventService EventService
}

func (v *VideoHasBeenTranscoded) Dispatch() error {

	data := map[string]interface{}{
		"video_title": v.VideoTitle,
	}

	jsonData, _ := json.Marshal(data)

	KAFKATOPIC := os.Getenv("KAFKATOPIC")

	err := v.EventService.Publish(jsonData, KAFKATOPIC)

	if err != nil {
		return err
	}
	return nil
}
