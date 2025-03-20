package kafka

import (
	"log"

	"context"

	"github.com/IBM/sarama"
)

type HandlerKafka interface {
	Start(ctx context.Context, brokers []string, group string, topics []string) error
	SetConsumer(consumer *ConsumerKafka)
	HandleNotificationCreated(message *sarama.ConsumerMessage)
}

type ConsumerKafka struct {
	Ready   chan bool
	handler HandlerKafka
}

func NewConsumer(handler HandlerKafka) *ConsumerKafka {
	return &ConsumerKafka{
		Ready:   make(chan bool),
		handler: handler,
	}
}

func (consumer *ConsumerKafka) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.Ready)
	return nil
}

func (consumer *ConsumerKafka) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *ConsumerKafka) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		switch message.Topic {
		case "notification_created":
			consumer.handler.HandleNotificationCreated(message)

		default:
			log.Printf("Получено сообщение из неизвестного топика %s: %s", message.Topic, string(message.Value))
		}

		// Отмечаем сообщение как обработанное
		session.MarkMessage(message, "")
	}
	return nil
}
