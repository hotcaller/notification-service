// internal/domains/notifications/service.go

package notifications

import (
	"context"
	"service/internal/domains/notifications/models"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ListNotifications(ctx context.Context) ([]models.Notification, error) {
	return s.repo.GetAllNotifications(ctx)
}

func (s *Service) GetNotificationByID(ctx context.Context, id int) (*models.Notification, error) {
	return s.repo.GetNotificationByID(ctx, id)
}

func (s *Service) SendNotification(ctx context.Context, notification models.Notification) error {
	return s.repo.SaveNotification(ctx, notification)
}
