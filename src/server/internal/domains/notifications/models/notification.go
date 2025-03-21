package models

import "time"

type NotificationType string

const (
    NotificationTypeNews      NotificationType = "news"
    NotificationTypeReminder  NotificationType = "reminder"
    NotificationTypeWarning   NotificationType = "warning"
    NotificationTypeImportant NotificationType = "important"
)

type Notification struct {
    ID        int             `json:"id" db:"id"`
    Header    string          `json:"header" db:"header"`
    Message   string          `json:"message" db:"message"`
    Type      NotificationType `json:"type" db:"type"`
    TargetID  int64           `json:"target_id" db:"target_id"`
    OrgToken  string          `json:"org_token" db:"org_token"`
    CreatedAt time.Time       `json:"created_at" db:"created_at"`
}