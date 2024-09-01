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

type DriverUsecase struct {
	db         *gorm.DB
	repository *repository.DriverRepository
}

func NewDriverUsecase(
	db *gorm.DB,
	repository *repository.DriverRepository,
) *DriverUsecase {
	return &DriverUsecase{
		db:         db,
		repository: repository,
	}
}

func (uc *DriverUsecase) Create(ctx context.Context, request *model.CreateDriverRequest) (*model.DriverResponse, error) {
	tx := uc.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	exists, err := uc.repository.CheckIfNIKOrPhoneNumberExists(tx, request.NIK, request.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, apperror.BadRequest(fmt.Errorf(`driver with NIK "%s" or phone number "%s" already exists`, request.NIK, request.PhoneNumber))
	}

	driver := &entity.Driver{
		Name:        request.Name,
		NIK:         request.NIK,
		PhoneNumber: request.PhoneNumber,
	}

	if err := uc.repository.Create(tx, driver); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, err
	}

	return model.ToDriverResponse(driver), nil
}

func (uc *DriverUsecase) Get(ctx context.Context, request *model.GetDriverRequest) (*model.DriverResponse, error) {
	tx := uc.db.WithContext(ctx).Begin(&sql.TxOptions{ReadOnly: true})
	defer tx.Rollback()

	driver, err := uc.repository.FindByID(tx, request.ID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, err
	}

	return model.ToDriverResponse(driver), nil
}

func (uc *DriverUsecase) GetList(ctx context.Context, request *model.GetListDriverRequest) ([]model.DriverResponse, int64, error) {
	tx := uc.db.WithContext(ctx).Begin(&sql.TxOptions{ReadOnly: true})
	defer tx.Rollback()

	drivers, total, err := uc.repository.Search(tx, request.NIK, request.Name, request.PhoneNumber, request.DailyCost, request.Page, request.Size)
	if err != nil {
		return nil, 0, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, 0, err
	}

	return model.ToDriversResponse(drivers), total, nil
}

func (uc *DriverUsecase) Update(ctx context.Context, request *model.UpdateDriverRequest) (*model.DriverResponse, error) {
	tx := uc.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	driver, err := uc.repository.FindByID(tx, request.ID)
	if err != nil {
		return nil, err
	}

	if request.NIK != nil {
		driver.NIK = *request.NIK
	}
	if request.Name != nil {
		driver.Name = *request.Name
	}
	if request.PhoneNumber != nil {
		driver.PhoneNumber = *request.PhoneNumber
	}
	if request.DailyCost != nil {
		driver.DailyCost = *request.DailyCost
	}

	if err := uc.repository.Update(tx, driver); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, err
	}

	return model.ToDriverResponse(driver), nil
}

func (uc *DriverUsecase) Delete(ctx context.Context, request *model.DeleteDriverRequest) error {
	tx := uc.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	driver, err := uc.repository.FindByID(tx, request.ID)
	if err != nil {
		return err
	}

	if err := uc.repository.Delete(tx, driver); err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return err
	}

	return nil
}
