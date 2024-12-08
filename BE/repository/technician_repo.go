package repository

import (
	"context"
	"errors"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/entity"
	errorUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/error"
	"gorm.io/gorm"
)

type (
	TechnicianRepository interface {
		FindByUserId(ctx context.Context, tx *gorm.DB, userID int64) (entity.Technician, errorUtils.CustomError)
		FindById(ctx context.Context, tx *gorm.DB, id int64) (entity.Technician, errorUtils.CustomError)
	}

	technicianRepository struct {
		db *gorm.DB
	}
)

func NewTechnicianRepository(db *gorm.DB) TechnicianRepository {
	return &technicianRepository{
		db: db,
	}
}

func (r *technicianRepository) FindByUserId(ctx context.Context, tx *gorm.DB, userID int64) (entity.Technician, errorUtils.CustomError) {
	if tx == nil {
		tx = r.db
	}

	var technician entity.Technician
	err := tx.WithContext(ctx).Where("user_id = ?", userID).First(&technician).Error

	if err != nil {
		return entity.Technician{}, errorUtils.ErrInternalServer
	}

	return technician, nil
}

func (r *technicianRepository) FindById(ctx context.Context, tx *gorm.DB, id int64) (entity.Technician, errorUtils.CustomError) {
	if tx == nil {
		tx = r.db
	}

	var technician entity.Technician
	err := tx.WithContext(ctx).Where("id = ?", id).First(&technician).Error

	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return entity.Technician{}, errorUtils.ErrTechnicianNotFound
		} else {
			return entity.Technician{}, errorUtils.ErrInternalServer
		}
	}

	return technician, nil
}
