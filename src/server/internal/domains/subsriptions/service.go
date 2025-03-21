package subscriptions

import (
	"context"
	"service/internal/domains/subscriptions/models"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetSubscriptionsByToken(ctx context.Context, token int64) ([]models.Subscription, error) {
	return s.repo.GetSubscriptionsByToken(ctx, token)
}