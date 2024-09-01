package entity

type DriverIncentive struct {
	ID        int     `gorm:"column:id;primaryKey"`
	BookingID int     `gorm:"column:booking_id"`
	Incentive float64 `gorm:"column:incentive"`
}

func (*DriverIncentive) TableName() string {
	return "driver_incentives"
}
