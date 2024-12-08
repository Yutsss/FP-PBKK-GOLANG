package repository

import (
	"context"
	"errors"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/dto"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/entity"
	errorUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/error"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	TicketRepository interface {
		ExtractContext(ctx context.Context) (*gorm.DB, errorUtils.CustomError)
		Create(ctx context.Context, tx *gorm.DB, data dto.CreateTicketRequest) (entity.Ticket, errorUtils.CustomError)
		FindAll(ctx context.Context, tx *gorm.DB) ([]entity.Ticket, errorUtils.CustomError)
		FindById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.Ticket, errorUtils.CustomError)
		FindByUserId(ctx context.Context, tx *gorm.DB, userID int64) ([]entity.Ticket, errorUtils.CustomError)
		UpdateById(ctx context.Context, tx *gorm.DB, id uuid.UUID, data entity.Ticket) (entity.Ticket, errorUtils.CustomError)
	}

	ticketRepository struct {
		db *gorm.DB
	}
)

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{
		db: db,
	}
}

func (r *ticketRepository) ExtractContext(ctx context.Context) (*gorm.DB, errorUtils.CustomError) {
	tx := r.db.WithContext(ctx)

	if tx == nil {
		return nil, errorUtils.ErrInternalServer
	}

	return tx, nil
}

func (r *ticketRepository) Create(ctx context.Context, tx *gorm.DB, data dto.CreateTicketRequest) (entity.Ticket, errorUtils.CustomError) {
	if tx == nil {
		tx = r.db
	}

	ticket := entity.Ticket{
		UserID:      data.UserID,
		Title:       data.Title,
		Description: data.Description,
		Category:    data.Category,
	}

	err := tx.WithContext(ctx).Create(&ticket).Error

	if err != nil {
		return entity.Ticket{}, errorUtils.ErrInternalServer
	}

	return ticket, nil
}

func (r *ticketRepository) FindAll(ctx context.Context, tx *gorm.DB) ([]entity.Ticket, errorUtils.CustomError) {
	if tx == nil {
		tx = r.db
	}

	var tickets []entity.Ticket

	err := tx.WithContext(ctx).Find(&tickets).Error

	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return []entity.Ticket{}, nil
		} else {
			return nil, errorUtils.ErrInternalServer
		}
	}

	return tickets, nil
}

func (r *ticketRepository) FindById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.Ticket, errorUtils.CustomError) {
	if tx == nil {
		tx = r.db
	}

	var ticket entity.Ticket

	err := tx.WithContext(ctx).Where("id = ?", id).First(&ticket).Error

	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return entity.Ticket{}, nil
		} else {
			return entity.Ticket{}, errorUtils.ErrInternalServer
		}
	}

	return ticket, nil
}

func (r *ticketRepository) FindByUserId(ctx context.Context, tx *gorm.DB, userID int64) ([]entity.Ticket, errorUtils.CustomError) {
	if tx == nil {
		tx = r.db
	}

	var tickets []entity.Ticket

	err := tx.WithContext(ctx).Where("user_id = ?", userID).Find(&tickets).Error

	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return []entity.Ticket{}, nil
		} else {
			return nil, errorUtils.ErrInternalServer
		}
	}

	return tickets, nil
}

func (r *ticketRepository) UpdateById(ctx context.Context, tx *gorm.DB, id uuid.UUID, data entity.Ticket) (entity.Ticket, errorUtils.CustomError) {
	if tx == nil {
		tx = r.db
	}

	var ticket entity.Ticket

	err := tx.WithContext(ctx).Where("id = ?", id).Updates(data).Error

	if err != nil {
		return entity.Ticket{}, errorUtils.ErrInternalServer
	}

	return ticket, nil
}
