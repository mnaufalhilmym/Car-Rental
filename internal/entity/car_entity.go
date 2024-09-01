package entity

import "gorm.io/gorm"

type Car struct {
	ID        int            `gorm:"column:id;primaryKey"`
	Name      string         `gorm:"column:name"`
	Stock     int            `gorm:"column:stock"`
	DailyRent float64        `gorm:"column:daily_rent"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (*Car) TableName() string {
	return "cars"
}
