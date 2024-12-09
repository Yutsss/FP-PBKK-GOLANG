package service

import (
	"context"
	"database/sql"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/constants"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/dto"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/repository"
	errorUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/error"
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/validation"
	"github.com/google/uuid"
)

type (
	TicketService interface {
		Create(ctx context.Context, data dto.CreateTicketRequest) (dto.CreateTicketResponse, errorUtils.CustomError)
		GetAll(ctx context.Context) (dto.GetAllTicketResponse, errorUtils.CustomError)
		GetById(ctx context.Context, data dto.GetTicketByIDRequest) (dto.GetTicketByIDResponse, errorUtils.CustomError)
		GetByUserId(ctx context.Context, data dto.GetTicketByUserIDRequest) (dto.GetTicketByUserIDResponse, errorUtils.CustomError)
		AssignById(ctx context.Context, data dto.AssignTicketByIdRequest) errorUtils.CustomError
		CloseById(ctx context.Context, data dto.CloseTicketByIdRequest) errorUtils.CustomError
	}

	ticketService struct {
		ticketRepo repository.TicketRepository
		adminRepo  repository.AdminRepository
		techRepo   repository.TechnicianRepository
		logRepo    repository.LogRepository
	}
)

func NewTicketService(ticketRepo repository.TicketRepository, adminRepo repository.AdminRepository, techRepo repository.TechnicianRepository, logRepo repository.LogRepository) TicketService {
	return &ticketService{
		ticketRepo: ticketRepo,
		adminRepo:  adminRepo,
		techRepo:   techRepo,
		logRepo:    logRepo,
	}
}

func (s *ticketService) Create(ctx context.Context, data dto.CreateTicketRequest) (dto.CreateTicketResponse, errorUtils.CustomError) {
	if err := validation.Validate(data); err != nil {
		return dto.CreateTicketResponse{}, err
	}

	ticket, err := s.ticketRepo.Create(ctx, nil, data)
	if err != nil {
		return dto.CreateTicketResponse{}, err
	}

	res := dto.CreateTicketResponse{
		ID:          ticket.ID,
		UserID:      ticket.UserID,
		Title:       ticket.Title,
		Description: ticket.Description,
		Category:    ticket.Category,
		Status:      ticket.Status,
	}

	return res, nil
}

func (s *ticketService) GetAll(ctx context.Context) (dto.GetAllTicketResponse, errorUtils.CustomError) {
	tickets, err := s.ticketRepo.FindAll(ctx, nil)
	if err != nil {
		return dto.GetAllTicketResponse{}, err
	}

	res := dto.GetAllTicketResponse{
		Tickets: make([]dto.TicketsResponse, 0),
	}

	for _, ticket := range tickets {
		res.Tickets = append(res.Tickets, dto.TicketsResponse{
			ID:           ticket.ID,
			UserID:       ticket.UserID,
			AdminID:      ticket.AdminID.Int64,
			LogID:        ticket.LogID.Int64,
			TechnicianID: ticket.TechnicianID.Int64,
			Title:        ticket.Title,
			Description:  ticket.Description,
			Category:     ticket.Category,
			Status:       ticket.Status,
		})
	}

	return res, nil
}

func (s *ticketService) GetById(ctx context.Context, data dto.GetTicketByIDRequest) (dto.GetTicketByIDResponse, errorUtils.CustomError) {
	if err := validation.Validate(data); err != nil {
		return dto.GetTicketByIDResponse{}, err
	}

	ticket, err := s.ticketRepo.FindById(ctx, nil, data.ID)
	if err != nil {
		return dto.GetTicketByIDResponse{}, err
	}

	if ticket.ID == uuid.Nil {
		return dto.GetTicketByIDResponse{}, errorUtils.ErrTicketNotFound
	}

	log, _ := s.logRepo.GetById(ctx, nil, ticket.LogID.Int64)

	if log.ID != 0 {
		res := dto.GetTicketByIDResponse{
			Ticket: dto.TicketResponse{
				ID:           ticket.ID,
				UserID:       ticket.UserID,
				AdminID:      ticket.AdminID.Int64,
				LogID:        ticket.LogID.Int64,
				TechnicianID: ticket.TechnicianID.Int64,
				Title:        ticket.Title,
				Description:  ticket.Description,
				Category:     ticket.Category,
				Status:       ticket.Status,
				Log: dto.LogResponse{
					ID:           log.ID,
					TechnicianID: log.TechnicianID,
					Activity:     log.Activity,
				},
			},
		}

		return res, nil
	}

	res := dto.GetTicketByIDResponse{
		Ticket: dto.TicketResponse{
			ID:           ticket.ID,
			UserID:       ticket.UserID,
			AdminID:      ticket.AdminID.Int64,
			LogID:        ticket.LogID.Int64,
			TechnicianID: ticket.TechnicianID.Int64,
			Title:        ticket.Title,
			Description:  ticket.Description,
			Category:     ticket.Category,
			Status:       ticket.Status,
		},
	}

	return res, nil
}

func (s *ticketService) GetByUserId(ctx context.Context, data dto.GetTicketByUserIDRequest) (dto.GetTicketByUserIDResponse, errorUtils.CustomError) {
	if err := validation.Validate(data); err != nil {
		return dto.GetTicketByUserIDResponse{}, err
	}

	tickets, err := s.ticketRepo.FindByUserId(ctx, nil, data.UserID)
	if err != nil {
		return dto.GetTicketByUserIDResponse{}, err
	}

	res := dto.GetTicketByUserIDResponse{
		Tickets: make([]dto.TicketsResponse, 0),
	}

	for _, ticket := range tickets {
		res.Tickets = append(res.Tickets, dto.TicketsResponse{
			ID:           ticket.ID,
			UserID:       ticket.UserID,
			AdminID:      ticket.AdminID.Int64,
			LogID:        ticket.LogID.Int64,
			TechnicianID: ticket.TechnicianID.Int64,
			Title:        ticket.Title,
			Description:  ticket.Description,
			Category:     ticket.Category,
			Status:       ticket.Status,
		})
	}

	return res, nil
}

func (s *ticketService) AssignById(ctx context.Context, data dto.AssignTicketByIdRequest) errorUtils.CustomError {
	if err := validation.Validate(data); err != nil {
		return err
	}

	tx, err := s.ticketRepo.ExtractContext(ctx)

	if err != nil {
		return err
	}

	tx.Begin()

	ticket, err := s.ticketRepo.FindById(ctx, tx, data.ID)

	if err != nil {
		tx.Rollback()
		return err
	}

	if ticket.ID == uuid.Nil {
		tx.Rollback()
		return errorUtils.ErrTicketNotFound
	}

	if ticket.Status != constants.ENUM_STATUS_OPEN {
		tx.Rollback()
		return errorUtils.ErrTicketAlreadyAssigned
	}

	admin, err := s.adminRepo.FindByUserId(ctx, tx, data.UserID)

	if err != nil {
		tx.Rollback()
		return err
	}

	technician, err := s.techRepo.FindById(ctx, tx, data.TechnicianID)

	if err != nil {
		tx.Rollback()
		return err
	}

	ticket.TechnicianID = sql.NullInt64{Int64: technician.ID, Valid: true}
	ticket.AdminID = sql.NullInt64{Int64: admin.ID, Valid: true}
	ticket.Status = constants.ENUM_STATUS_PENDING

	_, err = s.ticketRepo.UpdateById(ctx, tx, data.ID, ticket)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *ticketService) CloseById(ctx context.Context, data dto.CloseTicketByIdRequest) errorUtils.CustomError {
	if err := validation.Validate(data); err != nil {
		return err
	}

	tx, err := s.ticketRepo.ExtractContext(ctx)

	if err != nil {
		return err
	}

	tx.Begin()

	technician, err := s.techRepo.FindByUserId(ctx, tx, data.UserID)

	if err != nil {
		tx.Rollback()
		return err
	}

	data.TechnicianID = technician.ID

	log, err := s.logRepo.Create(ctx, tx, data)

	if err != nil {
		tx.Rollback()
		return err
	}

	ticket, err := s.ticketRepo.FindById(ctx, tx, data.TicketId)

	if err != nil {
		tx.Rollback()
		return err
	}

	if ticket.ID == uuid.Nil {
		tx.Rollback()
		return errorUtils.ErrTicketNotFound
	}

	if ticket.Status != constants.ENUM_STATUS_PENDING {
		tx.Rollback()
		return errorUtils.ErrTicketNotAssigned
	}

	if ticket.TechnicianID.Int64 != data.TechnicianID {
		tx.Rollback()
		return errorUtils.ErrNotAllowed
	}

	ticket.Status = constants.ENUM_STATUS_CLOSED
	ticket.LogID = sql.NullInt64{Int64: log.ID, Valid: true}

	_, err = s.ticketRepo.UpdateById(ctx, tx, data.TicketId, ticket)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
