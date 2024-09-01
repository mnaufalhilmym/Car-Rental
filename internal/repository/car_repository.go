package repository

import (
	"carrental/internal/entity"
	"carrental/internal/util"
	"fmt"
	"strconv"
	"strings"

	"github.com/mnaufalhilmym/gotracing"
	"gorm.io/gorm"
)

type CarRepository struct {
	repository[entity.Car]
}

func NewCarRepository(db *gorm.DB) *CarRepository {
	if err := db.AutoMigrate(&entity.Car{}); err != nil {
		if err.Error() != fmt.Sprintf(`ERROR: relation "%s" already exists (SQLSTATE 42P07)`, (&entity.Car{}).TableName()) {
			panic(fmt.Errorf("failed to migrate entity: %w", err))
		}
	}

	return &CarRepository{}
}

func (r *CarRepository) SearchCars(
	db *gorm.DB,
	name string,
	stock string,
	dailyRent string,
	page int,
	size int,
) ([]entity.Car, int64, error) {
	var cars []entity.Car
	var total int64

	offset := 0
	if page > 0 {
		offset = (page - 1) * size
	}

	filter := r.filterCars(name, stock, dailyRent)

	if err := db.Scopes(filter).Offset(offset).Limit(size).Find(&cars).Error; err != nil {
		gotracing.Error("Failed to find entities from database", err)
		return nil, 0, err
	}

	if err := db.Model(&entity.Car{}).Scopes(filter).Count(&total).Error; err != nil {
		gotracing.Error("Failed to count entities from database", err)
		return nil, 0, err
	}

	return cars, total, nil
}

func (*CarRepository) filterCars(
	nameFilter string,
	stockFilter string,
	dailyRentFilter string,
) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if nameFilter != "" {
			fName := "%" + nameFilter + "%"
			tx = tx.Where("name ILIKE ?", fName)
		}

		if stockFilter != "" {
			filter := strings.Split(nameFilter, "=")
			if len(filter) == 2 {
				stock, err := strconv.Atoi(filter[1])
				if err == nil {
					if operator := util.ParseComparisonFilter(filter[0]); operator != "" {
						tx = tx.Where("stock "+operator+" ?", stock)
					}
				}
			}
		}

		if dailyRentFilter != "" {
			filter := strings.Split(dailyRentFilter, "=")
			if len(filter) == 2 {
				dailyRent, err := strconv.ParseFloat(filter[1], 64)
				if err == nil {
					if operator := util.ParseComparisonFilter(filter[0]); operator != "" {
						tx = tx.Where("daily_rent "+operator+" ?", dailyRent)
					}
				}
			}
		}

		return tx
	}
}
