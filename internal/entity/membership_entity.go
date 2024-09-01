package entity

import "gorm.io/gorm"

type Membership struct {
	ID             int            `gorm:"column:id;primaryKey"`
	MembershipName string         `gorm:"column:membership_name"`
	Discount       float64        `gorm:"column:discount"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (*Membership) TableName() string {
	return "memberships"
}
