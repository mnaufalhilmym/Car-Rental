package repository

import (
	"carrental/internal/entity"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mnaufalhilmym/gotracing"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	repository[entity.Customer]
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	if err := db.Migrator().CreateTable(&entity.Customer{}); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); !ok || pgErr.Code != "42P07" {
			panic(fmt.Errorf("failed to migrate entity: %w", err))
		}
	}

	return &CustomerRepository{}
}

func (r *CustomerRepository) CheckIfNIKOrPhoneNumberExists(db *gorm.DB, nik string, phoneNumber string) (bool, error) {
	var count int64
	if err := db.Model(&entity.Customer{}).Where("nik = ? OR phone_number = ?", nik, phoneNumber).Count(&count).Error; err != nil {
		gotracing.Error("Failed to find entity from database", err)
		return false, err
	}
	return count > 0, nil
}

func (r *CustomerRepository) SearchCustomers(db *gorm.DB, nik string, name string, phoneNumber string, page int, size int) ([]entity.Customer, int64, error) {
	var customers []entity.Customer
	var total int64

	offset := 0
	if page > 0 {
		offset = (page - 1) * size
	}

	filter := r.filterCustomers(nik, name, phoneNumber)

	if err := db.Scopes(filter).Offset(offset).Limit(size).Find(&customers).Error; err != nil {
		gotracing.Error("Failed to find entities from database", err)
		return nil, 0, err
	}

	if err := db.Model(&entity.Customer{}).Scopes(filter).Count(&total).Error; err != nil {
		gotracing.Error("Failed to count entities from database", err)
		return nil, 0, err
	}

	return customers, total, nil
}

func (*CustomerRepository) filterCustomers(nik string, name string, phoneNumber string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if nik != "" {
			fNIK := "%" + nik + "%"
			tx = tx.Where("nik LIKE ?", fNIK)
		}

		if name != "" {
			fName := "%" + name + "%"
			tx = tx.Where("name ILIKE ?", fName)
		}

		if phoneNumber != "" {
			fPhoneNumber := "%" + phoneNumber + "%"
			tx = tx.Where("phone_number LIKE ?", fPhoneNumber)
		}

		return tx
	}
}
