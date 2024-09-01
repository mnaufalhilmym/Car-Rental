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

type MembershipUsecase struct {
	db         *gorm.DB
	repository *repository.MembershipRepository
}

func NewMembershipUsecase(
	db *gorm.DB,
	repository *repository.MembershipRepository,
) *MembershipUsecase {
	return &MembershipUsecase{
		db:         db,
		repository: repository,
	}
}

func (uc *MembershipUsecase) Create(ctx context.Context, request *model.CreateMembershipRequest) (*model.MembershipResponse, error) {
	tx := uc.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	exists, err := uc.repository.CheckIfMembershipNameInsensitiveExists(tx, request.MembershipName)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, apperror.BadRequest(fmt.Errorf(`membership with name "%s" already exists`, request.MembershipName))
	}

	membership := &entity.Membership{
		MembershipName: request.MembershipName,
		Discount:       request.Discount,
	}

	if err := uc.repository.Create(tx, membership); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, err
	}

	return model.ToMembershipResponse(membership), nil
}

func (uc *MembershipUsecase) GetAll(ctx context.Context) ([]model.MembershipResponse, error) {
	tx := uc.db.WithContext(ctx).Begin(&sql.TxOptions{ReadOnly: true})
	defer tx.Rollback()

	memberships, err := uc.repository.FindAll(tx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		gotracing.Error("Failed to commit transaction", err)
		return nil, err
	}

	return model.ToMembershipsResponse(memberships), nil
}

func (uc *MembershipUsecase) Delete(ctx context.Context, request *model.DeleteMembershipRequest) error {
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
