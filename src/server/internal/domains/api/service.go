package api

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"service/internal/infrastructure/config"
	"sort"
	"strings"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) ValidateTelegramData(data url.Values) bool {
	checkHash := data.Get("hash")
	if checkHash == "" {
		return false
	}

	dataCopy := url.Values{}
	for k, v := range data {
		if k != "hash" {
			dataCopy[k] = v
		}
	}

	keys := make([]string, 0, len(dataCopy))
	for k := range dataCopy {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var dataStrings []string
	for _, key := range keys {
		value := dataCopy.Get(key)
		dataStrings = append(dataStrings, key+"="+value)
	}
	dataCheckString := strings.Join(dataStrings, "\n")

	cfg := config.GetTelegram()

	secretKey := sha256.Sum256([]byte(cfg.BotToken))
	h := hmac.New(sha256.New, secretKey[:])
	h.Write([]byte(dataCheckString))
	calculatedHash := hex.EncodeToString(h.Sum(nil))

	return calculatedHash == checkHash
}

//func (s*Service) ProcessTelegramLogin(ctx context.Context, userData url.Values) error {*
//	// Extract user data
//	idStr := userData.Get("id")
//	if idStr == "" {
//		return fmt.Errorf("отсутствует идентификатор пользователя")
//	}
//	id, err := strconv.ParseInt(idStr, 10, 64)
//	if err != nil {
//		return fmt.Errorf("некорректный идентификатор пользователя: %w", err)
//	}
//
//	firstName := userData.Get("first_name")
//	lastName := userData.Get("last_name")
//	username := userData.Get("username")
//	photoURL := userData.Get("photo_url")
//	authDateStr := userData.Get("auth_date")
//	authDate, err := strconv.ParseInt(authDateStr, 10, 64)
//	if err != nil {
//		return fmt.Errorf("некорректная дата аутентификации: %w", err)
//	}
//
//	user := TelegramUser{
//		ID:        id,
//		FirstName: firstName,
//		LastName:  lastName,
//		Username:  username,
//		PhotoURL:  photoURL,
//		AuthDate:  authDate,
//	}
//
//	err = s.repo.SaveTelegramUser(ctx, user)
//	if err != nil {
//		return fmt.Errorf("не удалось сохранить данные пользователя: %w", err)
//	}
//	return nil
//}

func (s *Service) GenerateQRCode(ctx context.Context, patientID string, token string) ([]byte, error) {
	qrCodeData, err := s.repo.GenerateQRCodeData(ctx, patientID, token)
	if err != nil {
		return nil, fmt.Errorf("не удалось сгенерировать QR-код: %w", err)
	}
	return qrCodeData, nil
}
