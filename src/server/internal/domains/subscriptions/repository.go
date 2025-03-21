package subscriptions

import (
	"context"
	"fmt"
	"github.com/Arlandaren/pgxWrappy/pkg/postgres"
	"service/internal/domains/subscriptions/models"
)

type Repository struct {
	db *postgres.Wrapper
}

func NewRepository(db *postgres.Wrapper) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetSubscriptionsByToken(ctx context.Context, token string) ([]models.Subscription, error) {
	query := `
        SELECT id, telegram_id, token, patient_id
        FROM subscriptions
        WHERE token = $1
    `
	var subscriptions []models.Subscription

	if err := r.db.Select(ctx, &subscriptions, query, token); err != nil {
		return nil, fmt.Errorf("failed to get subscriptions: %w", err)
	}
	return subscriptions, nil
}