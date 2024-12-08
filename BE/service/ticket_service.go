package service

import (
	"context"
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
	}

	ticketService struct {
		ticketRepo repository.TicketRepository
	}
)

func NewTicketService(ticketRepo repository.TicketRepository) TicketService {
	return &ticketService{
		ticketRepo: ticketRepo,
	}
}

func (s *ticketService) Create(ctx context.Context, data dto.CreateTicketRequest) (dto.CreateTicketResponse, errorUtils.CustomError) {
	if err := validation.Validate(data); err != nil {
		return dto.CreateTicketResponse{}, err
	}

	ticket, err := s.ticketRepo.Create(ctx, nil, data)
	if err != nil {
		return dto.CreateTicketResponse{}, errorUtils.ErrInternalServer
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
		return dto.GetAllTicketResponse{}, errorUtils.ErrInternalServer
	}

	res := dto.GetAllTicketResponse{
		Tickets: make([]dto.TicketResponse, 0),
	}

	for _, ticket := range tickets {
		res.Tickets = append(res.Tickets, dto.TicketResponse{
			ID:           ticket.ID,
			UserID:       ticket.UserID,
			AdminID:      ticket.AdminID,
			LogID:        ticket.LogID,
			TechnicianID: ticket.TechnicianID,
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
		return dto.GetTicketByIDResponse{}, errorUtils.ErrInternalServer
	}

	if ticket.ID == uuid.Nil {
		return dto.GetTicketByIDResponse{}, errorUtils.ErrTicketNotFound
	}

	res := dto.GetTicketByIDResponse{
		Ticket: dto.TicketResponse{
			ID:           ticket.ID,
			UserID:       ticket.UserID,
			AdminID:      ticket.AdminID,
			LogID:        ticket.LogID,
			TechnicianID: ticket.TechnicianID,
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

	tickets, err := s.ticketRepo.FindByUserID(ctx, nil, data.UserID)
	if err != nil {
		return dto.GetTicketByUserIDResponse{}, errorUtils.ErrInternalServer
	}

	res := dto.GetTicketByUserIDResponse{
		Tickets: make([]dto.TicketResponse, 0),
	}

	for _, ticket := range tickets {
		res.Tickets = append(res.Tickets, dto.TicketResponse{
			ID:           ticket.ID,
			UserID:       ticket.UserID,
			AdminID:      ticket.AdminID,
			LogID:        ticket.LogID,
			TechnicianID: ticket.TechnicianID,
			Title:        ticket.Title,
			Description:  ticket.Description,
			Category:     ticket.Category,
			Status:       ticket.Status,
		})
	}

	return res, nil
}
