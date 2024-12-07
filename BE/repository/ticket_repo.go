package repository

import (
	"context"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/dto"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/entity"
	errorUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/error"
	"gorm.io/gorm"
)

type (
	TicketRepository interface {
		Create(ctx context.Context, tx *gorm.DB, data dto.CreateTicketRequest) (entity.Ticket, errorUtils.CustomError)
		FindAll(ctx context.Context, tx *gorm.DB) ([]entity.Ticket, errorUtils.CustomError)
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
		return nil, errorUtils.ErrInternalServer
	}

	return tickets, nil
}
