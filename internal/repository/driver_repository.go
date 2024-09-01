package repository

import (
	"carrental/internal/entity"
	"carrental/internal/util"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mnaufalhilmym/gotracing"
	"gorm.io/gorm"
)

type DriverRepository struct {
	repository[entity.Driver]
}

func NewDriverRepository(db *gorm.DB) *DriverRepository {
	if err := db.Migrator().CreateTable(&entity.Driver{}); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); !ok || pgErr.Code != "42P07" {
			panic(fmt.Errorf("failed to migrate entity: %w", err))
		}
	}

	return &DriverRepository{}
}

func (r *DriverRepository) CheckIfNIKOrPhoneNumberExists(db *gorm.DB, nik string, phoneNumber string) (bool, error) {
	var count int64
	if err := db.Model(&entity.Driver{}).Where("nik = ? OR phone_number = ?", nik, phoneNumber).Count(&count).Error; err != nil {
		gotracing.Error("Failed to find entity from database", err)
		return false, err
	}
	return count > 0, nil
}

func (r *DriverRepository) Search(db *gorm.DB, nik string, name string, phoneNumber string, dailyCost string, page int, size int) ([]entity.Driver, int64, error) {
	var drivers []entity.Driver
	var total int64

	offset := 0
	if page > 0 {
		offset = (page - 1) * size
	}

	filter := r.filter(nik, name, phoneNumber, dailyCost)

	if err := db.Scopes(filter).Offset(offset).Limit(size).Find(&drivers).Error; err != nil {
		gotracing.Error("Failed to find entities from database", err)
		return nil, 0, err
	}

	if err := db.Model(&entity.Driver{}).Scopes(filter).Count(&total).Error; err != nil {
		gotracing.Error("Failed to count entities from database", err)
		return nil, 0, err
	}

	return drivers, total, nil
}

func (*DriverRepository) filter(nik string, name string, phoneNumber string, dailyCostFilter string) func(tx *gorm.DB) *gorm.DB {
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

		if dailyCostFilter != "" {
			filter := strings.Split(dailyCostFilter, "=")
			if len(filter) == 2 {
				dailyCost, err := strconv.ParseFloat(filter[1], 64)
				if err == nil {
					if operator := util.ParseComparisonFilter(filter[0]); operator != "" {
						tx = tx.Where("daily_cost "+operator+" ?", dailyCost)
					}
				}
			}
		}

		return tx
	}
}
