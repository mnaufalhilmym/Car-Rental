package entity

import "gorm.io/gorm"

type Customer struct {
	ID          int            `gorm:"column:id;primaryKey"`
	Name        string         `gorm:"column:name"`
	NIK         string         `gorm:"column:nik"`
	PhoneNumber string         `gorm:"column:phone_number"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (*Customer) TableName() string {
	return "customers"
}
