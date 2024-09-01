package repository

import (
	"carrental/internal/entity"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mnaufalhilmym/gotracing"
	"gorm.io/gorm"
)

type BookingTypeRepository struct {
	repository[entity.BookingType]
}

func NewBookingTypeRepository(db *gorm.DB) *BookingTypeRepository {
	if err := db.Migrator().CreateTable(&entity.BookingType{}); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); !ok || pgErr.Code != "42P07" {
			panic(fmt.Errorf("failed to migrate entity: %w", err))
		}
	}

	return &BookingTypeRepository{}
}

func (r *BookingTypeRepository) CheckIfBookingTypeInsensitiveExists(db *gorm.DB, bookingType string) (bool, error) {
	var count int64
	if err := db.Model(&entity.BookingType{}).Where("LOWER(booking_type) = LOWER(?)", bookingType).Count(&count).Error; err != nil {
		gotracing.Error("Failed to find entity from database", err)
		return false, err
	}
	return count > 0, nil
}
