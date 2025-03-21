package application

import (
	"github.com/Arlandaren/pgxWrappy/pkg/postgres"
	"service/internal/domains/api"
	"service/internal/domains/feedback"
	"service/internal/domains/notifications"
	"service/internal/domains/subscriptions"  
	"service/internal/infrastructure/storage/minio"
	"service/internal/infrastructure/storage/redis"
)

type Repository struct {
	Notification *notifications.Repository
	Api          *api.Repository
	Subscription *subscriptions.Repository
	Feedback     *feedback.Repository
}

func NewRepository(db *postgres.Wrapper, rdb *redis.RDB, s3 *minio.Minio) *Repository {
	return &Repository{
		Notification: notifications.NewRepository(db),
		Api:          api.NewRepository(db, rdb, s3),
		Subscription: subscriptions.NewRepository(db),
		Feedback:     feedback.NewRepository(db),
	}
}