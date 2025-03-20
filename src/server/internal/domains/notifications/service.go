// internal/domains/notifications/service.go

package notifications

import (
	"context"
	"encoding/json"
	"service/internal/domains/notifications/models"
	"service/internal/infrastructure/kafka"
)

type Service struct {
	repo     *Repository
	Producer *kafka.Producer
}

func NewService(repo *Repository, producer *kafka.Producer) *Service {
	return &Service{repo: repo, Producer: producer}
}

func (s *Service) ListNotifications(ctx context.Context, userID int64) ([]models.Notification, error) {
	return s.repo.GetNotificationsByUserID(ctx, userID)
}

func (s *Service) GetNotificationByID(ctx context.Context, id int64, userID int64) (*models.Notification, error) {
	return s.repo.GetNotificationByIDAndUserID(ctx, id, userID)
}

func (s *Service) SendNotification(ctx context.Context, notification models.Notification) error {
	data, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	// Отправляем JSON как сообщение в Kafka
	err = s.Producer.ProduceMessage("notification_created", data)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ProcessNotifications(ctx context.Context, notification *models.Notification) error {
	return s.repo.SaveNotification(ctx, notification)
}
