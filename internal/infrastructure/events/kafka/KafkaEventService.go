package kafka

import (
	"context"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type EventService struct {
	writer *kafka.Writer
}

func NewKafkaEventService(address []string) *EventService {
	fmt.Sprintf("starting a writer to kafka with address %s", address[0])
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

	fmt.Sprintf("publishing message to topic %s with a message %v", destination, jsonData)
	err := service.writer.WriteMessages(context.Background(),
		kafka.Message{
			Topic: destination,
			Value: jsonData,
		},
	)
	fmt.Println("message published")
	if err != nil {
		return err
	}

	return nil
}
