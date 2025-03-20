package application

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"

	"service/internal/domains/notifications/models"
	"service/internal/infrastructure/kafka"
)

type AppConsumer struct {
	s        *Service
	consumer *kafka.ConsumerKafka
}

func NewAppConsumer(svc *Service, consumer *kafka.ConsumerKafka) *AppConsumer {
	return &AppConsumer{
		s:        svc,
		consumer: consumer,
	}
}

func (c *AppConsumer) SetConsumer(consumer *kafka.ConsumerKafka) {
	c.consumer = consumer
}

func (c *AppConsumer) Start(ctx context.Context, brokers []string, group string, topics []string) error {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{
		sarama.NewBalanceStrategyRange(),
	}

	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	client, err := sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		return fmt.Errorf("error creating consumer group: %v", err)
	}

	// Создаем канал для ошибок
	consumerErrors := make(chan error, 1)

	// Запускаем в отдельной горутине цикл потребления
	go func() {
		defer close(consumerErrors)
		for {
			// Проверяем контекст на отмену
			if ctx.Err() != nil {
				return
			}
			// Начинаем потребление
			err := client.Consume(ctx, topics, c.consumer)
			if err != nil {
				consumerErrors <- err
				return
			}
			// Сбрасываем готовность после ребалансировки
			c.consumer.Ready = make(chan bool)
		}
	}()

	// Ожидаем, пока потребитель будет готов
	<-c.consumer.Ready
	log.Println("ConsumerKafka is ready and waiting for messages...")

	// Ожидаем завершения работы или возникновения ошибки
	select {
	case <-ctx.Done():
		log.Println("Context canceled. Exiting ConsumerKafka...")
		return client.Close()
	case err := <-consumerErrors:
		return fmt.Errorf("error in ConsumerKafka: %v", err)
	}
}

func (c *AppConsumer) HandleNotificationCreated(message *sarama.ConsumerMessage) {
	var data models.Notification

	err := json.Unmarshal(message.Value, &data)
	if err != nil {
		log.Printf("decode error notification_created: %v", err)
		return
	}

	ctx := context.Background()

	err = c.s.Notification.ProcessNotifications(ctx, &data)

	if err != nil {
		log.Printf("handle error booking_created: %v", err)
	}
	log.Printf("Получено сообщение из топика %s: %s", message.Topic, string(message.Value))
}
