package application

import (
	"service/internal/domains/api"
	"service/internal/domains/feedback"
	"service/internal/domains/notifications"
	"service/internal/domains/subscriptions" 
	"service/internal/infrastructure/kafka"
)

type Service struct {
	Notification *notifications.Service
	Api          *api.Service
	Subscription *subscriptions.Service
	Feedback     *feedback.Service
	Producer     *kafka.Producer
}

func NewService(repo *Repository, producer *kafka.Producer) *Service {
	return &Service{
		Notification: notifications.NewService(repo.Notification, producer),
		Api:          api.NewService(repo.Api),
		Subscription: subscriptions.NewService(repo.Subscription),
		Feedback:     feedback.NewService(repo.Feedback),
		Producer:     producer,
	}
}