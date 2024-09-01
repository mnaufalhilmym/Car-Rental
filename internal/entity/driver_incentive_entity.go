package entity

type DriverIncentive struct {
	ID        int     `gorm:"column:id;primaryKey"`
	BookingID int     `gorm:"column:booking_id;unique"`
	Incentive float64 `gorm:"column:incentive"`

	Booking Booking `gorm:"foreignKey:booking_id;references:id"`
}

func (*DriverIncentive) TableName() string {
	return "driver_incentives"
}
