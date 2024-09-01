package entity

type Car struct {
	ID        int     `gorm:"column:id;primaryKey"`
	Name      string  `gorm:"column:name;unique"`
	Stock     int     `gorm:"column:stock"`
	DailyRent float64 `gorm:"column:daily_rent"`
}

func (*Car) TableName() string {
	return "cars"
}
