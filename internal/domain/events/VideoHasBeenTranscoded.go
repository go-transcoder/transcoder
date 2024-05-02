package events

import (
	"encoding/json"
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

	err := v.EventService.Publish(jsonData, "video-transcoded")

	if err != nil {
		return err
	}
	return nil
}
