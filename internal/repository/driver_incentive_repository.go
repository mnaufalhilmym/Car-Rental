package repository

import (
	"carrental/internal/entity"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type DriverIncentiveRepository struct {
	repository[entity.DriverIncentive]
}

func NewDriverIncentiveRepository(db *gorm.DB) *DriverIncentiveRepository {
	if err := db.Migrator().CreateTable(&entity.DriverIncentive{}); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); !ok || pgErr.Code != "42P07" {
			panic(fmt.Errorf("failed to migrate entity: %w", err))
		}
	}

	return &DriverIncentiveRepository{}
}
