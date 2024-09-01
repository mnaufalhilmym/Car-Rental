package entity

import "gorm.io/gorm"

type Customer struct {
	ID           int            `gorm:"column:id;primaryKey"`
	Name         string         `gorm:"column:name"`
	NIK          string         `gorm:"column:nik"`
	PhoneNumber  string         `gorm:"column:phone_number"`
	MembershipID *int           `gorm:"column:membership_id"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`

	Membership *Membership `gorm:"foreignKey:membership_id;references:id"`
}

func (*Customer) TableName() string {
	return "customers"
}
