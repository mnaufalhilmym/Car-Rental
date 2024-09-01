package entity

import "time"

type Booking struct {
	ID         int       `gorm:"column:id;primaryKey"`
	CustomerID int       `gorm:"column:customer_id"`
	CarID      int       `gorm:"column:car_id"`
	StartRent  time.Time `gorm:"column:start_rent"`
	EndRent    time.Time `gorm:"column:end_rent"`
	TotalCost  float64   `gorm:"column:total_cost"`
	Finished   bool      `gorm:"column:finished"`
	Customer   Customer  `gorm:"foreignKey:customer_id;references:id"`
	Car        Car       `gorm:"foreignKey:car_id;references:id"`
}

func (*Booking) TableName() string {
	return "bookings"
}
