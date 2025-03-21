package models

import "time"

type Feedback struct {
    ID        int64      `json:"id"`
    Header    string     `json:"header"`
    Content   string     `json:"content"`
    Answer    *string    `json:"answer,omitempty"`
    UserID    int64      `json:"user_id"`
    CreatedAt time.Time  `json:"created_at"`
    AnsweredAt *time.Time `json:"answered_at,omitempty"`
}