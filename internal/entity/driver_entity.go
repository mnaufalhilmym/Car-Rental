package entity

type Driver struct {
	ID          int     `gorm:"column:id;primaryKey"`
	Name        string  `gorm:"column:name"`
	NIK         string  `gorm:"column:nik"`
	PhoneNumber string  `gorm:"column:phone_number"`
	DailyCost   float64 `gorm:"column:daily_cost"`
}

func (*Driver) TableName() string {
	return "drivers"
}
