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
    ID        int             `json:"id"`
    Header    string          `json:"header"`  
    Message   string          `json:"message"`
    Type      NotificationType `json:"type"`
    TargetID  int64           `json:"target_id"`
    OrgToken  string          `json:"org_token"`
    CreatedAt time.Time       `json:"created_at"`
}