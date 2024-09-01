package usecase

import (
	"carrental/internal/entity"
	"carrental/internal/model"
	"carrental/internal/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/mnaufalhilmym/gotracing"
	"gorm.io/gorm"
)

type CustomerUsecase struct {
	db         *gorm.DB
	validator  *validator.Validate
	repository *repository.CustomerRepository
}

func NewCustomerUsecase(
	db *gorm.DB,
	validator *validator.Validate,
	repository *repository.CustomerRepository,
) *CustomerUsecase {
	return &CustomerUsecase{
		db:         db,
		validator:  validator,
		repository: repository,
	}
}

func (uc *CustomerUsecase) Create(ctx context.Context, request *model.CreateCustomerRequest) (*model.CustomerResponse, error) {
	tx := uc.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := uc.validator.Struct(request); err != nil {
		gotracing.Error("Failed to validate request", err)
		return nil, err
	}

	customer := &entity.Customer{
		Name:        request.Name,
		NIK:         request.NIK,
		PhoneNumber: request.PhoneNumber,
	}

	if err := uc.repository.Create(tx, customer); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, err
	}

	return model.ToCustomerResponse(customer), nil
}

func (uc *CustomerUsecase) Get(ctx context.Context, request *model.GetCustomerRequest) (*model.CustomerResponse, error) {
	tx := uc.db.WithContext(ctx).Begin(&sql.TxOptions{ReadOnly: true})
	defer tx.Rollback()

	if err := uc.validator.Struct(request); err != nil {
		gotracing.Error("Failed to validate request", err)
		return nil, err
	}

	customer, err := uc.repository.FindByID(tx, request.ID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, err
	}

	return model.ToCustomerResponse(customer), nil
}

func (uc *CustomerUsecase) GetList(ctx context.Context, request *model.GetListCustomerRequest) ([]model.CustomerResponse, int64, error) {
	tx := uc.db.WithContext(ctx).Begin(&sql.TxOptions{ReadOnly: true})
	defer tx.Rollback()

	if err := uc.validator.Struct(request); err != nil {
		gotracing.Error("Failed to validate request", err)
		return nil, 0, err
	}

	customers, total, err := uc.repository.SearchCustomers(tx, request.NIK, request.Name, request.PhoneNumber, request.Page, request.Size)
	if err != nil {
		return nil, 0, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, 0, err
	}

	return model.ToCustomersResponse(customers), total, nil
}

func (uc *CustomerUsecase) Update(ctx context.Context, request *model.UpdateCustomerRequest) (*model.CustomerResponse, error) {
	tx := uc.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := uc.validator.Struct(request); err != nil {
		gotracing.Error("Failed to validate request", err)
		return nil, err
	}

	customer, err := uc.repository.FindByID(tx, request.ID)
	if err != nil {
		return nil, err
	}

	if request.NIK != nil {
		customer.NIK = *request.NIK
	}
	if request.Name != nil {
		customer.Name = *request.Name
	}
	if request.PhoneNumber != nil {
		customer.PhoneNumber = *request.PhoneNumber
	}

	if err := uc.repository.Update(tx, customer); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, err
	}

	return model.ToCustomerResponse(customer), nil
}

func (uc *CustomerUsecase) Delete(ctx context.Context, request *model.DeleteCustomerRequest) error {
	tx := uc.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := uc.validator.Struct(request); err != nil {
		gotracing.Error("Failed to validate request", err)
		return err
	}

	customer, err := uc.repository.FindByID(tx, request.ID)
	if err != nil {
		return err
	}

	if err := uc.repository.Delete(tx, customer); err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return err
	}

	return nil
}
