package entity

import "gorm.io/gorm"

type Driver struct {
	ID          int            `gorm:"column:id;primaryKey"`
	Name        string         `gorm:"column:name"`
	NIK         string         `gorm:"column:nik"`
	PhoneNumber string         `gorm:"column:phone_number"`
	DailyCost   float64        `gorm:"column:daily_cost"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (*Driver) TableName() string {
	return "drivers"
}
