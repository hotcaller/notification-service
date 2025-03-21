package notifications

import (
    "context"
    "database/sql"
    "fmt"
    "time"

    _ "github.com/lib/pq" // PostgreSQL driver
    "service/internal/domains/notifications/models"
)

type Repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
    return &Repository{
        db: db,
    }
}

func (r *Repository) GetAllNotifications(ctx context.Context) ([]models.Notification, error) {
    query := `
        SELECT id, header, message, type, target_id, org_token, created_at
        FROM notifications
        ORDER BY created_at DESC
    `
    
    rows, err := r.db.QueryContext(ctx, query)
    if err != nil {
        return nil, fmt.Errorf("не удалось получить уведомления: %w", err)
    }
    defer rows.Close()
    
    var notifications []models.Notification
    
    for rows.Next() {
        var n models.Notification
        var createdAt time.Time
        
        err := rows.Scan(
            &n.ID,
            &n.Header,
            &n.Message,
            &n.Type,
            &n.TargetID,
            &n.OrgToken,
            &createdAt,
        )
        if err != nil {
            return nil, fmt.Errorf("не удалось считать уведомление: %w", err)
        }
        
        n.CreatedAt = createdAt
        notifications = append(notifications, n)
    }
    
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("ошибка при итерации по результатам: %w", err)
    }
    
    return notifications, nil
}

func (r *Repository) GetNotificationsByUserID(ctx context.Context, patientID int64) ([]models.Notification, error) {
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
    
    // Convert patientID to string for the query
    patientIDStr := fmt.Sprintf("%d", patientID)
    
    rows, err := r.db.QueryContext(ctx, query, patientID, patientIDStr)
    if err != nil {
        return nil, fmt.Errorf("не удалось получить уведомления: %w", err)
    }
    defer rows.Close()
    
    var notifications []models.Notification
    
    for rows.Next() {
        var n models.Notification
        var createdAt time.Time
        
        err := rows.Scan(
            &n.ID,
            &n.Header,
            &n.Message,
            &n.Type,
            &n.TargetID,
            &n.OrgToken,
            &createdAt,
        )
        if err != nil {
            return nil, fmt.Errorf("не удалось считать уведомление: %w", err)
        }
        
        n.CreatedAt = createdAt
        notifications = append(notifications, n)
    }
    
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("ошибка при итерации по результатам: %w", err)
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
    
    // Convert patientID to string for the query
    patientIDStr := fmt.Sprintf("%d", patientID)
    
    var n models.Notification
    var createdAt time.Time
    
    err := r.db.QueryRowContext(ctx, query, notificationID, patientID, patientIDStr).Scan(
        &n.ID,
        &n.Header,
        &n.Message,
        &n.Type,
        &n.TargetID,
        &n.OrgToken,
        &createdAt,
    )
    
    if err == sql.ErrNoRows {
        return nil, nil // No notification found
    }
    
    if err != nil {
        return nil, fmt.Errorf("не удалось получить уведомление: %w", err)
    }
    
    n.CreatedAt = createdAt
    return &n, nil
}

func (r *Repository) SaveNotification(ctx context.Context, notification *models.Notification) error {
    query := `
        INSERT INTO notifications (header, message, type, target_id, org_token)
        VALUES ($1, $2, $3, $4, $5)
    `
    
    _, err := r.db.ExecContext(ctx, query, 
        notification.Header, 
        notification.Message, 
        string(notification.Type), // Convert NotificationType to string
        notification.TargetID, 
        notification.OrgToken)
    
    if err != nil {
        return fmt.Errorf("не удалось сохранить уведомление: %w", err)
    }
    
    return nil
}