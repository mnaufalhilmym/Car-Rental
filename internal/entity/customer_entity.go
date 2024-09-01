package entity

type Customer struct {
	ID          int    `gorm:"column:id;primaryKey"`
	Name        string `gorm:"column:name"`
	NIK         string `gorm:"column:nik;unique"`
	PhoneNumber string `gorm:"column:phone_number;unique"`
}

func (*Customer) TableName() string {
	return "customers"
}
