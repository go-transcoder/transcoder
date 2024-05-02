package events

type EventService interface {
	Publish(jsonData []byte, destination string) error
}
