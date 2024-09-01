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

type CustomerUsecase struct {
	db                   *gorm.DB
	repository           *repository.CustomerRepository
	membershipRepository *repository.MembershipRepository
}

func NewCustomerUsecase(
	db *gorm.DB,
	repository *repository.CustomerRepository,
	membershipRepository *repository.MembershipRepository,
) *CustomerUsecase {
	return &CustomerUsecase{
		db:                   db,
		repository:           repository,
		membershipRepository: membershipRepository,
	}
}

func (uc *CustomerUsecase) Create(ctx context.Context, request *model.CreateCustomerRequest) (*model.CustomerResponse, error) {
	tx := uc.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	exists, err := uc.repository.CheckIfNIKOrPhoneNumberExists(tx, request.NIK, request.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, apperror.BadRequest(fmt.Errorf(`customer with NIK "%s" or phone number "%s" already exists`, request.NIK, request.PhoneNumber))
	}

	var membership *entity.Membership
	if request.MembershipID != nil {
		_membership, err := uc.membershipRepository.FindByID(tx, *request.MembershipID)
		if err != nil {
			return nil, err
		}
		membership = _membership
	}

	customer := &entity.Customer{
		Name:         request.Name,
		NIK:          request.NIK,
		PhoneNumber:  request.PhoneNumber,
		MembershipID: request.MembershipID,

		Membership: membership,
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

	customer, err := uc.repository.FindByID(tx, request.ID)
	if err != nil {
		return nil, err
	}

	if customer.MembershipID != nil {
		if err := uc.repository.LoadMembership(tx, customer); err != nil {
			return nil, err
		}
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

	customers, total, err := uc.repository.Search(tx, request.NIK, request.Name, request.PhoneNumber, request.Page, request.Size)
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
	if request.MembershipID != nil {
		membership, err := uc.membershipRepository.FindByID(tx, *request.MembershipID)
		if err != nil {
			return nil, err
		}
		customer.MembershipID = &membership.ID
		customer.Membership = membership
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
