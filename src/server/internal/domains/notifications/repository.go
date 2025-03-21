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
        SELECT id, header, message, type, target_id, org_token, created_at
        FROM notifications
        ORDER BY created_at DESC
    `
    if err := r.db.Select(ctx, &notifications, query); err != nil {
        return nil, fmt.Errorf("не удалось получить уведомления: %w", err)
    }
    return notifications, nil
}

func (r *Repository) GetNotificationsByUserID(ctx context.Context, userID int64) ([]models.Notification, error) {
    // This query assumes:
    // 1. Notifications are sent to target_id which is the patient_id
    // 2. The userID parameter is actually the patient_id 
    // 3. Broadcast notifications (target_id = 0) for organizations are also included
    query := `
    SELECT n.id, n.header, n.message, n.type, n.target_id, n.org_token, n.created_at
    FROM notifications n
    WHERE n.target_id = $1
    OR (
        n.target_id = 0 AND n.org_token = (
            SELECT org_token 
            FROM patients 
            WHERE id = $1
        )
    )
    ORDER BY n.created_at DESC
    `
    var notifications []models.Notification

    if err := r.db.Select(ctx, &notifications, query, userID); err != nil {
        return nil, err
    }
    return notifications, nil
}

func (r *Repository) GetNotificationByIDAndUserID(ctx context.Context, id int64, userID int64) (*models.Notification, error) {
    query := `
    SELECT n.id, n.header, n.message, n.type, n.target_id, n.org_token, n.created_at
    FROM notifications n
    WHERE n.id = $1 AND (
        n.target_id = $2
        OR (
            n.target_id = 0 AND n.org_token = (
                SELECT org_token 
                FROM patients 
                WHERE id = $2
            )
        )
    )
    `
    var n models.Notification

    if err := r.db.Get(ctx, &n, query, id, userID); err != nil {
        return nil, err
    }
    return &n, nil
}

func (r *Repository) SaveNotification(ctx context.Context, notification *models.Notification) error {
    query := `
        INSERT INTO notifications (header, message, type, target_id, org_token)
        VALUES ($1, $2, $3, $4, $5)
    `
    _, err := r.db.Exec(ctx, query, 
        notification.Header, 
        notification.Message, 
        notification.Type, 
        notification.TargetID, 
        notification.OrgToken)
    if err != nil {
        return fmt.Errorf("не удалось сохранить уведомление: %w", err)
    }
    return nil
}