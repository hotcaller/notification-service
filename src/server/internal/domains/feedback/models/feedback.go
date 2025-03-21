package models

type Feedback struct {
	ID         int64   `json:"id" db:"id"`
	Header     string  `json:"header" db:"header"`
	Content    string  `json:"content" db:"content"`
	Answer     *string `json:"answer,omitempty" db:"answer"`
	UserID     int64   `json:"user_id" db:"user_id"`
	CreatedAt  string  `json:"created_at" db:"created_at"`
	AnsweredAt *string `json:"answered_at,omitempty" db:"answered_at"`
}
