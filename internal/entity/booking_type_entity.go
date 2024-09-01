package entity

import "gorm.io/gorm"

type BookingType struct {
	ID          int            `gorm:"column:id;primaryKey"`
	BookingType string         `gorm:"column:booking_type"`
	Description string         `gorm:"column:description"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (*BookingType) TableName() string {
	return "booking_types"
}
