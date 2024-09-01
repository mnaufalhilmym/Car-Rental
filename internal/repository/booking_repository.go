package repository

import (
	"carrental/internal/entity"
	"carrental/internal/util"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mnaufalhilmym/gotracing"
	"gorm.io/gorm"
)

type BookingRepository struct {
	repository[entity.Booking]
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	if err := db.Migrator().CreateTable(&entity.Booking{}); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); !ok || pgErr.Code != "42P07" {
			panic(fmt.Errorf("failed to migrate entity: %w", err))
		}
	}

	return &BookingRepository{}
}

func (r *BookingRepository) LoadCustomer(db *gorm.DB, booking *entity.Booking) error {
	if err := db.Where("id = ?", booking.CustomerID).First(&booking.Customer).Error; err != nil {
		gotracing.Error("Failed to find entity from database", err)
		return err
	}
	return nil
}

func (r *BookingRepository) LoadCar(db *gorm.DB, booking *entity.Booking) error {
	if err := db.Where("id = ?", booking.CarID).First(&booking.Car).Error; err != nil {
		gotracing.Error("Failed to find entity from database", err)
		return err
	}
	return nil
}

func (*BookingRepository) FindByIDPreload(db *gorm.DB, id int) (*entity.Booking, error) {
	var entity *entity.Booking
	if err := db.Joins("Customer").Joins("Car").Joins("BookingType").Joins("LEFT JOIN drivers ON drivers.id = driver_id").Where("bookings.id = ?", id).First(&entity).Error; err != nil {
		gotracing.Error("Failed to find entity from database", err)
		return nil, err
	}
	return entity, nil
}

func (r *BookingRepository) SearchPreload(
	db *gorm.DB,
	customerID int,
	carID int,
	startRent string,
	endRent string,
	offsetTime time.Duration,
	totalCost string,
	finished *bool,
	page int,
	size int,
) ([]entity.Booking, int64, error) {
	var bookings []entity.Booking
	var total int64

	offset := 0
	if page > 0 {
		offset = (page - 1) * size
	}

	filter := r.filter(customerID, carID, startRent, endRent, offsetTime, totalCost, finished)

	if err := db.Joins("Customer").Joins("Car").Scopes(filter).Offset(offset).Limit(size).Find(&bookings).Error; err != nil {
		gotracing.Error("Failed to find entities from database", err)
		return nil, 0, err
	}

	if err := db.Model(&entity.Booking{}).Scopes(filter).Count(&total).Error; err != nil {
		gotracing.Error("Failed to count entities from database", err)
		return nil, 0, err
	}

	return bookings, total, nil

}

func (*BookingRepository) filter(
	customerID int,
	carID int,
	startRentFilter string,
	endRentFilter string,
	offset time.Duration,
	totalCostFilter string,
	finished *bool,
) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		var startRent time.Time
		var startRentOperator string
		var endRent time.Time
		var endRentOperator string

		if customerID > 0 {
			tx = tx.Where("customer_id = ?", customerID)
		}

		if carID > 0 {
			tx = tx.Where("car_id = ?", carID)
		}

		if startRentFilter != "" {
			filter := strings.Split(startRentFilter, "=")

			filterValue := filter[0]
			if len(filter) == 2 {
				filterValue = filter[1]
				if operator := util.ParseComparisonFilter(filter[0]); operator != "" {
					startRentOperator = operator
				}
			}

			_startRent, err := time.Parse("2006-01-02", filterValue)
			if err == nil {
				startRent = _startRent.Add(offset)
			}
		}

		if endRentFilter != "" {
			filter := strings.Split(endRentFilter, "=")

			filterValue := filter[0]
			if len(filter) == 2 {
				filterValue = filter[1]
				if operator := util.ParseComparisonFilter(filter[0]); operator != "" {
					endRentOperator = operator
				}
			}

			_endRent, err := time.Parse("2006-01-02", filterValue)
			if err == nil {
				endRent = _endRent.Add(offset)
			}
		}

		if !startRent.IsZero() && !endRent.IsZero() {
			tx = tx.Where("(start_rent < ? AND end_rent >= ?) OR (start_rent >= ? AND start_rent < ?)", startRent, startRent, startRent, endRent)
		} else if !startRent.IsZero() && startRentOperator != "" {
			tx = tx.Where("start_rent "+startRentOperator+" ?", startRent)
		} else if !endRent.IsZero() && endRentOperator != "" {
			tx = tx.Where("end_rent "+endRentOperator+" ?", endRentOperator)
		}

		if totalCostFilter != "" {
			filter := strings.Split(totalCostFilter, "=")
			if len(filter) == 2 {
				totalCost, err := strconv.ParseFloat(filter[1], 64)
				if err == nil {
					if operator := util.ParseComparisonFilter(filter[0]); operator != "" {
						tx = tx.Where("total_cost "+operator+" ?", totalCost)
					}
				}
			}
		}

		if finished != nil {
			tx = tx.Where("finished = ?", *finished)
		}

		return tx
	}
}
