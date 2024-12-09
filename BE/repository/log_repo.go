package repository

import (
	"context"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/dto"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/entity"
	errorUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/error"
	"gorm.io/gorm"
)

type (
	LogRepository interface {
		Create(ctx context.Context, tx *gorm.DB, data dto.CloseTicketByIdRequest) (entity.Log, errorUtils.CustomError)
		GetById(ctx context.Context, tx *gorm.DB, Id int64) (entity.Log, errorUtils.CustomError)
	}

	logRepository struct {
		db *gorm.DB
	}
)

func NewLogRepository(db *gorm.DB) LogRepository {
	return &logRepository{
		db: db,
	}
}

func (r *logRepository) Create(ctx context.Context, tx *gorm.DB, data dto.CloseTicketByIdRequest) (entity.Log, errorUtils.CustomError) {
	if tx == nil {
		tx = r.db
	}

	log := entity.Log{
		TechnicianID: data.TechnicianID,
		Activity:     data.Activity,
	}

	err := tx.WithContext(ctx).Create(&log).Error

	if err != nil {
		return entity.Log{}, errorUtils.ErrInternalServer
	}

	return log, nil
}

func (r *logRepository) GetById(ctx context.Context, tx *gorm.DB, Id int64) (entity.Log, errorUtils.CustomError) {
	if tx == nil {
		tx = r.db
	}

	var log entity.Log
	err := tx.WithContext(ctx).Where("id = ?", Id).Find(&log).Error

	if err != nil {
		return entity.Log{}, errorUtils.ErrInternalServer
	}

	return log, nil
}
