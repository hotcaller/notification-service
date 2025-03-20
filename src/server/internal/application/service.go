package application

import (
	"service/internal/domains/api"
	"service/internal/domains/notifications"
	"service/internal/infrastructure/kafka"
)

type Service struct {
	Notification *notifications.Service
	Api          *api.Service
	Producer     *kafka.Producer
}

func NewService(repo *Repository, producer *kafka.Producer) *Service {
	return &Service{
		Notification: notifications.NewService(repo.Notification, producer),
		Api:          api.NewService(repo.Api),
	}
}
