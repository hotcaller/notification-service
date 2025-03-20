package models

import "time"

type Notification struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	TargetID  int64     `json:"target_id"`
	OrgToken  string    `json:"org_token"`
	CreatedAt time.Time `json:"created_at"`
}
