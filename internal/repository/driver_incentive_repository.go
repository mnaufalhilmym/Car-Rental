package repository

import (
	"carrental/internal/entity"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mnaufalhilmym/gotracing"
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

func (*DriverIncentiveRepository) FindByBookingID(db *gorm.DB, bookingID int) (*entity.DriverIncentive, error) {
	var entity *entity.DriverIncentive
	if err := db.Where("booking_id = ?", bookingID).First(&entity).Error; err != nil {
		gotracing.Error("Failed to find entity from database", err)
		return nil, err
	}
	return entity, nil
}
