package repository

import (
	"context"
	"errors"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/entity"
	errorUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/error"
	"gorm.io/gorm"
)

type (
	AdminRepository interface {
		FindByUserId(ctx context.Context, tx *gorm.DB, userID int64) (entity.Admin, errorUtils.CustomError)
	}

	adminRepository struct {
		db *gorm.DB
	}
)

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{
		db: db,
	}
}

func (r *adminRepository) FindByUserId(ctx context.Context, tx *gorm.DB, userID int64) (entity.Admin, errorUtils.CustomError) {
	if tx == nil {
		tx = r.db
	}

	var admin entity.Admin
	err := tx.WithContext(ctx).Where("user_id = ?", userID).First(&admin).Error

	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return entity.Admin{}, errorUtils.ErrNotAllowed
		} else {
			return entity.Admin{}, errorUtils.ErrInternalServer
		}
	}

	return admin, nil
}
