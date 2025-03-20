package notifications

import (
	"context"
	"fmt"

	"github.com/Arlandaren/pgxWrappy/pkg/postgres"
	"service/internal/domains/notifications/models"
)

type Repository struct {
	db *postgres.Wrapper
}

func NewRepository(db *postgres.Wrapper) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAllNotifications(ctx context.Context) ([]models.Notification, error) {
	var notifications []models.Notification
	query := `
        SELECT id, message, target_id, org_token, created_at
        FROM notifications
        ORDER BY created_at DESC
    `
	if err := r.db.Select(ctx, notifications, query); err != nil {
		return nil, fmt.Errorf("не удалось получить уведомления: %w", err)
	}
	return notifications, nil
}

func (r *Repository) GetNotificationsByUserID(ctx context.Context, userID string) ([]models.Notification, error) {
	query := `SELECT id, message,target_id, org_token, created_at FROM notifications WHERE target_id = $1`
	var notifications []models.Notification

	if err := r.db.Select(ctx, &notifications, query, userID); err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *Repository) GetNotificationByIDAndUserID(ctx context.Context, id int, userID string) (*models.Notification, error) {
	query := `SELECT id, message,target_id, org_token, created_at FROM notifications WHERE id = $1 AND target_id = $2`
	var n models.Notification

	if err := r.db.Get(ctx, &n, query, id, userID); err != nil {
		return nil, err
	}
	return &n, nil
}

func (r *Repository) SaveNotification(ctx context.Context, notification models.Notification) error {
	query := `
        INSERT INTO notifications (message, target_id, org_token)
        VALUES ($1, $2, $3)
    `
	_, err := r.db.Exec(ctx, query, notification.Message, notification.TargetID, notification.OrgToken)
	if err != nil {
		return fmt.Errorf("не удалось сохранить уведомление: %w", err)
	}
	return nil
}
