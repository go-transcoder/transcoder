package events

import (
	"encoding/json"
)

type VideoHasBeenTranscoded struct {
	VideoTitle   string
	Topic        string
	EventService EventService
}

func (v *VideoHasBeenTranscoded) Dispatch() error {

	data := map[string]interface{}{
		"video_title": v.VideoTitle,
	}

	jsonData, _ := json.Marshal(data)

	err := v.EventService.Publish(jsonData, v.Topic)

	if err != nil {
		return err
	}
	return nil
}
