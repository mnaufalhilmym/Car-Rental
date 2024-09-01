package usecase

import (
	"carrental/internal/entity"
	apperror "carrental/internal/error"
	"carrental/internal/model"
	"carrental/internal/repository"
	"context"
	"database/sql"
	"fmt"

	"github.com/mnaufalhilmym/gotracing"
	"gorm.io/gorm"
)

type BookingTypeUsecase struct {
	db         *gorm.DB
	repository *repository.BookingTypeRepository
}

func NewBookingTypeUsecase(
	db *gorm.DB,
	repository *repository.BookingTypeRepository,
) *BookingTypeUsecase {
	return &BookingTypeUsecase{
		db:         db,
		repository: repository,
	}
}

func (uc *BookingTypeUsecase) Create(ctx context.Context, request *model.CreateBookingTypeRequest) (*model.BookingTypeResponse, error) {
	tx := uc.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	exists, err := uc.repository.CheckIfBookingTypeInsensitiveExists(tx, request.BookingType)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, apperror.BadRequest(fmt.Errorf(`booking type with type "%s" already exists`, request.BookingType))
	}

	bookingType := &entity.BookingType{
		BookingType: request.BookingType,
		Description: request.Description,
	}

	if err := uc.repository.Create(tx, bookingType); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, err
	}

	return model.ToBookingTypeResponse(bookingType), nil
}

func (uc *BookingTypeUsecase) GetAll(ctx context.Context) ([]model.BookingTypeResponse, error) {
	tx := uc.db.WithContext(ctx).Begin(&sql.TxOptions{ReadOnly: true})
	defer tx.Rollback()

	bookingTypes, err := uc.repository.FindAll(tx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, err
	}

	return model.ToBookingTypesResponse(bookingTypes), nil
}

func (uc *BookingTypeUsecase) Delete(ctx context.Context, request *model.DeleteBookingTypeRequest) error {
	tx := uc.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	car, err := uc.repository.FindByID(tx, request.ID)
	if err != nil {
		return err
	}

	if err := uc.repository.Delete(tx, car); err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return err
	}

	return nil
}
