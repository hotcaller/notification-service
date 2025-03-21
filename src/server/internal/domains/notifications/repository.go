package notifications

import (
    "context"
    "fmt"
		"time"
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

func (r *Repository) GetNotificationsByUserID(ctx context.Context, patientID int64) ([]models.Notification, error) {
    // This query gets notifications where:
    // 1. The target_id matches the patient_id directly
    // 2. OR it's a broadcast notification (target_id = 0) for the organization
    query := `
    SELECT n.id, n.header, n.message, n.type, n.target_id, n.org_token, n.created_at
    FROM notifications n
    WHERE n.target_id = $1
    OR (n.target_id = 0 AND n.org_token IN (
        SELECT s.token::text 
        FROM subscriptions s 
        WHERE s.patient_id = $2
    ))
    ORDER BY n.created_at DESC
    `
    var notifications []models.Notification

    // Convert patientID to string for the query
    patientIDStr := fmt.Sprintf("%d", patientID)
    
    if err := r.db.Select(ctx, &notifications, query, patientID, patientIDStr); err != nil {
        return nil, err
    }
    return notifications, nil
}

func (r *Repository) GetNotificationByIDAndUserID(ctx context.Context, notificationID int64, patientID int64) (*models.Notification, error) {
    query := `
    SELECT n.id, n.header, n.message, n.type, n.target_id, n.org_token, n.created_at
    FROM notifications n
    WHERE n.id = $1 AND (
        n.target_id = $2
        OR (n.target_id = 0 AND n.org_token IN (
            SELECT s.token::text 
            FROM subscriptions s 
            WHERE s.patient_id = $3
        ))
    )
    `
    var n models.Notification

    // Convert patientID to string for the query
    patientIDStr := fmt.Sprintf("%d", patientID)
    
    if err := r.db.Get(ctx, &n, query, notificationID, patientID, patientIDStr); err != nil {
        return nil, err
    }
    return &n, nil
}

func (r *Repository) SaveNotification(ctx context.Context, notification *models.Notification) error {
	notification.CreatedAt = time.Now()
	
	createdAtStr := notification.CreatedAt.Format(time.RFC3339)
	
	query := `
			INSERT INTO notifications (header, message, type, target_id, org_token, created_at)
			VALUES ($1, $2, $3, $4, $5, $6)
	`
	
	_, err := r.db.Exec(ctx, query, 
			notification.Header, 
			notification.Message, 
			notification.Type, 
			notification.TargetID, 
			notification.OrgToken,
			createdAtStr)
	
	if err != nil {
			return fmt.Errorf("не удалось сохранить уведомление: %w", err)
	}
	
	return nil
}