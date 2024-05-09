package kafka

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
)

type EventService struct {
	writer *kafka.Writer
}

func NewKafkaEventService(address []string) *EventService {
	return &EventService{
		writer: &kafka.Writer{
			Addr: kafka.TCP(address[0]), // TODO: put all the addresses
		},
	}
}

func (service *EventService) Publish(jsonData []byte, destination string) error {

	if service.writer == nil {
		return errors.New("make sure that you use the init method first")
	}

	err := service.writer.WriteMessages(context.Background(),
		kafka.Message{
			Topic: destination,
			Value: jsonData,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
