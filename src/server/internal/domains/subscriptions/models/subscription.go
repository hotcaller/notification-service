package models

type Subscription struct {
    ID         int    `json:"id" db:"id"`
    TelegramID string `json:"telegram_id" db:"telegram_id"`
    Token      string `json:"token" db:"token"`  
    PatientID  string `json:"patient_id" db:"patient_id"`
}