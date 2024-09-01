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

type CarUsecase struct {
	db         *gorm.DB
	repository *repository.CarRepository
}

func NewCarUsecase(
	db *gorm.DB,
	repository *repository.CarRepository,
) *CarUsecase {
	return &CarUsecase{
		db:         db,
		repository: repository,
	}
}

func (uc *CarUsecase) Create(ctx context.Context, request *model.CreateCarRequest) (*model.CarResponse, error) {
	tx := uc.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	exists, err := uc.repository.CheckIfNameInsensitiveExists(tx, request.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, apperror.BadRequest(fmt.Errorf(`car with name "%s" already exists`, request.Name))
	}

	car := &entity.Car{
		Name:      request.Name,
		Stock:     request.Stock,
		DailyRent: request.DailyRent,
	}

	if err := uc.repository.Create(tx, car); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, err
	}

	return model.ToCarResponse(car), nil
}

func (uc *CarUsecase) Get(ctx context.Context, request *model.GetCarRequest) (*model.CarResponse, error) {
	tx := uc.db.WithContext(ctx).Begin(&sql.TxOptions{ReadOnly: true})
	defer tx.Rollback()

	car, err := uc.repository.FindByID(tx, request.ID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, err
	}

	return model.ToCarResponse(car), nil
}

func (uc *CarUsecase) GetList(ctx context.Context, request *model.GetListCarRequest) ([]model.CarResponse, int64, error) {
	tx := uc.db.WithContext(ctx).Begin(&sql.TxOptions{ReadOnly: true})
	defer tx.Rollback()

	cars, total, err := uc.repository.SearchCars(tx, request.Name, request.Stock, request.DailyRent, request.Page, request.Size)
	if err != nil {
		return nil, 0, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, 0, err
	}

	return model.ToCarsResponse(cars), total, nil
}

func (uc *CarUsecase) Update(ctx context.Context, request *model.UpdateCarRequest) (*model.CarResponse, error) {
	tx := uc.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	car, err := uc.repository.FindByID(tx, request.ID)
	if err != nil {
		return nil, err
	}

	if request.Name != nil {
		car.Name = *request.Name
	}
	if request.Stock != nil {
		car.Stock = *request.Stock
	}
	if request.DailyRent != nil {
		car.DailyRent = *request.DailyRent
	}

	if err := uc.repository.Update(tx, car); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, err
	}

	return model.ToCarResponse(car), nil
}

func (uc *CarUsecase) Delete(ctx context.Context, request *model.DeleteCarRequest) error {
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
