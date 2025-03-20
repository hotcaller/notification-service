package application

import (
	"service/internal/domains/api"
	"service/internal/domains/notifications"
)

type Service struct {
	Notification *notifications.Service
	Api          *api.Service
}

func NewService(repo *Repository) *Service {
	return &Service{
		Notification: notifications.NewService(repo.Notification),
		Api:          api.NewService(repo.Api),
	}
}
