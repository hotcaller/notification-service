package models

type Subscription struct {
	ID        int       `json:"id"`
	TelegramID   string    `json:"telegram_id"`
	Token  int64     `json:"token"`
	PatientID  string    `json:"patient_id"`
}
