package api

import (
	"context"
	"fmt"

	"github.com/Arlandaren/pgxWrappy/pkg/postgres"
	"service/internal/infrastructure/storage/minio"
	"service/internal/infrastructure/storage/redis"

	"github.com/skip2/go-qrcode"
)

type Repository struct {
	db  *postgres.Wrapper
	rdb *redis.RDB
	s3  *minio.Minio
}

func NewRepository(db *postgres.Wrapper, rdb *redis.RDB, s3 *minio.Minio) *Repository {
	return &Repository{
		db:  db,
		rdb: rdb,
		s3:  s3,
	}
}

func (r *Repository) SaveTelegramUser(ctx context.Context, userData string) error {

	return nil
}

func (r *Repository) GenerateQRCodeData(ctx context.Context, patientID string, token string) ([]byte, error) {
	data := fmt.Sprintf("https://t.me/@ZabMedicalBot?start=%s|%s", patientID, token)

	qrCodeData, err := qrcode.Encode(data, qrcode.Medium, 256)
	if err != nil {
		return nil, fmt.Errorf("не удалось сгенерировать QR-код: %w", err)
	}

	// Сохраняем QR-код в MinIO
	// bucketName := "qrcodes"
	// objectName := fmt.Sprintf("%s.png", patientID)
	// err = r.s3.PutObject(ctx, bucketName, objectName, qrCodeData)
	// if err != nil {
	//  return nil, fmt.Errorf("не удалось сохранить QR-код в MinIO: %w", err)
	// }

	return qrCodeData, nil
}
