package entity

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID              int            `gorm:"column:id;primaryKey"`
	CustomerID      int            `gorm:"column:customer_id"`
	CarID           int            `gorm:"column:car_id"`
	StartRent       time.Time      `gorm:"column:start_rent"`
	EndRent         time.Time      `gorm:"column:end_rent"`
	TotalCost       float64        `gorm:"column:total_cost"` // days * daily_rent
	Finished        bool           `gorm:"column:finished"`
	Discount        float64        `gorm:"column:discount"` // total_cost * memberships.discount
	BookingTypeID   *int           `gorm:"column:booking_type_id"`
	DriverID        *int           `gorm:"column:driver_id"`
	TotalDriverCost float64        `gorm:"column:total_driver_cost"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at"`

	Customer        Customer         `gorm:"foreignKey:customer_id;references:id"`
	Car             Car              `gorm:"foreignKey:car_id;references:id"`
	BookingType     *BookingType     `gorm:"foreignKey:booking_type_id;references:id"`
	Driver          *Driver          `gorm:"foreignKey:driver_id;references:id"`
	DriverIncentive *DriverIncentive `gorm:"foreignKey:id;references:booking_id"`
}

func (*Booking) TableName() string {
	return "bookings"
}
