package repository

import (
	"carrental/internal/entity"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mnaufalhilmym/gotracing"
	"gorm.io/gorm"
)

type MembershipRepository struct {
	repository[entity.Membership]
}

func NewMembershipRepository(db *gorm.DB) *MembershipRepository {
	if err := db.Migrator().CreateTable(&entity.Membership{}); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); !ok || pgErr.Code != "42P07" {
			panic(fmt.Errorf("failed to migrate entity: %w", err))
		}
	}

	return &MembershipRepository{}
}

func (r *MembershipRepository) CheckIfMembershipNameInsensitiveExists(db *gorm.DB, membershipName string) (bool, error) {
	var count int64
	if err := db.Model(&entity.Membership{}).Where("LOWER(membership_name) = LOWER(?)", membershipName).Count(&count).Error; err != nil {
		gotracing.Error("Failed to find entity from database", err)
		return false, err
	}
	return count > 0, nil
}
