package entity

import "gorm.io/gorm"

type DriverIncentive struct {
	ID        int            `gorm:"column:id;primaryKey"`
	BookingID int            `gorm:"column:booking_id;unique"`
	Incentive float64        `gorm:"column:incentive"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	Booking Booking `gorm:"foreignKey:booking_id;references:id"`
}

func (*DriverIncentive) TableName() string {
	return "driver_incentives"
}
