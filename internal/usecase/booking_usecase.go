package usecase

import (
	"carrental/internal/entity"
	"carrental/internal/model"
	"carrental/internal/repository"
	"carrental/internal/util"
	"context"
	"database/sql"

	"github.com/mnaufalhilmym/gotracing"
	"gorm.io/gorm"
)

type BookingUsecase struct {
	db                 *gorm.DB
	repository         *repository.BookingRepository
	customerRepository *repository.CustomerRepository
	carRepository      *repository.CarRepository
}

func NewBookingUsecase(
	db *gorm.DB,
	repository *repository.BookingRepository,
	customerRepository *repository.CustomerRepository,
	carRepository *repository.CarRepository,
) *BookingUsecase {
	return &BookingUsecase{
		db:                 db,
		repository:         repository,
		customerRepository: customerRepository,
		carRepository:      carRepository,
	}
}

func (uc *BookingUsecase) Create(ctx context.Context, request *model.CreateBookingRequest) (*model.BookingResponse, error) {
	tx := uc.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	customer, err := uc.customerRepository.FindByID(tx, request.CustomerID)
	if err != nil {
		return nil, err
	}

	car, err := uc.carRepository.FindByID(tx, request.CarID)
	if err != nil {
		return nil, err
	}

	booking := &entity.Booking{
		CustomerID: customer.ID,
		CarID:      car.ID,
		StartRent:  request.StartRent,
		EndRent:    request.EndRent,
		TotalCost:  request.TotalCost,
		Finished:   request.Finished,
		Customer:   *customer,
		Car:        *car,
	}

	if err := uc.repository.Create(tx, booking); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, err
	}

	return model.ToBookingResponse(booking), nil
}

func (uc *BookingUsecase) Get(ctx context.Context, request *model.GetBookingRequest) (*model.BookingResponse, error) {
	tx := uc.db.WithContext(ctx).Begin(&sql.TxOptions{ReadOnly: true})
	defer tx.Rollback()

	booking, err := uc.repository.FindByIDPreload(tx, request.ID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, err
	}

	return model.ToBookingResponse(booking), nil
}

func (uc *BookingUsecase) GetList(ctx context.Context, request *model.GetListBookingRequest) ([]model.BookingResponse, int64, error) {
	tx := uc.db.WithContext(ctx).Begin(&sql.TxOptions{ReadOnly: true})
	defer tx.Rollback()

	offsetTime, err := util.ParseTimezone(request.Timezone)
	if err != nil {
		return nil, 0, err
	}

	bookings, total, err := uc.repository.SearchPreload(
		tx,
		request.CustomerID,
		request.CarID,
		request.StartRent,
		request.EndRent,
		offsetTime,
		request.TotalCost,
		request.Finished,
		request.Page,
		request.Size,
	)
	if err != nil {
		return nil, 0, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, 0, err
	}

	return model.ToBookingsResponse(bookings), total, nil
}

func (uc *BookingUsecase) Update(ctx context.Context, request *model.UpdateBookingRequest) (*model.BookingResponse, error) {
	tx := uc.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	booking, err := uc.repository.FindByIDPreload(tx, request.ID)
	if err != nil {
		return nil, err
	}

	if request.CustomerID != nil {
		customer, err := uc.customerRepository.FindByID(tx, *request.CustomerID)
		if err != nil {
			return nil, err
		}
		booking.CustomerID = customer.ID
	}
	if request.CarID != nil {
		car, err := uc.carRepository.FindByID(tx, *request.CarID)
		if err != nil {
			return nil, err
		}
		booking.CarID = car.ID
	}
	if request.StartRent != nil {
		booking.StartRent = *request.StartRent
	}
	if request.EndRent != nil {
		booking.EndRent = *request.EndRent
	}
	if request.TotalCost != nil {
		booking.TotalCost = *request.TotalCost
	}
	if request.Finished != nil {
		booking.Finished = *request.Finished
	}

	if err := uc.repository.Update(tx, booking); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, err
	}

	return model.ToBookingResponse(booking), nil
}

func (uc *BookingUsecase) Delete(ctx context.Context, request *model.DeleteBookingRequest) error {
	tx := uc.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	booking, err := uc.repository.FindByID(tx, request.ID)
	if err != nil {
		return err
	}

	if err := uc.repository.Delete(tx, booking); err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return err
	}

	return nil
}
