package api

import (
	"github.com/Arlandaren/pgxWrappy/pkg/postgres"
	"service/internal/infrastructure/storage/minio"
	"service/internal/infrastructure/storage/redis"
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
