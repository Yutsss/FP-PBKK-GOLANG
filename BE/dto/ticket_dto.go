package dto

import (
	"github.com/google/uuid"
)

type (
	TicketResponse struct {
		ID           uuid.UUID   `json:"id"`
		UserID       int64       `json:"user_id"`
		AdminID      int64       `json:"admin_id"`
		LogID        int64       `json:"log_id"`
		TechnicianID int64       `json:"technician_id"`
		Title        string      `json:"title"`
		Description  string      `json:"description"`
		Category     string      `json:"category"`
		Status       string      `json:"status"`
		Log          LogResponse `json:"log,omitempty"`
	}

	TicketsResponse struct {
		ID           uuid.UUID `json:"id"`
		UserID       int64     `json:"user_id"`
		AdminID      int64     `json:"admin_id"`
		LogID        int64     `json:"log_id"`
		TechnicianID int64     `json:"technician_id"`
		Title        string    `json:"title"`
		Description  string    `json:"description"`
		Category     string    `json:"category"`
		Status       string    `json:"status"`
	}

	CreateTicketRequest struct {
		UserID      int64  `json:"user_id" validate:"required"`
		Title       string `json:"title" validate:"required,min=1,max=255"`
		Description string `json:"description" validate:"required,min=1"`
		Category    string `json:"category" validate:"required,min=1,max=15"`
	}

	CreateTicketResponse struct {
		ID          uuid.UUID `json:"id"`
		UserID      int64     `json:"user_id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Category    string    `json:"category"`
		Status      string    `json:"status"`
	}

	GetTicketByIDRequest struct {
		ID uuid.UUID `json:"id" validate:"required"`
	}

	GetTicketByIDResponse struct {
		Ticket TicketResponse `json:"ticket"`
	}

	GetAllTicketResponse struct {
		Tickets []TicketsResponse `json:"tickets"`
	}

	GetTicketByUserIDRequest struct {
		UserID int64 `json:"user_id" validate:"required"`
	}

	GetTicketByUserIDResponse struct {
		Tickets []TicketsResponse `json:"tickets"`
	}

	UpdateTicketData struct {
		Title       string `json:"title" validate:"required,min=1,max=255"`
		Description string `json:"description" validate:"required,min=1"`
		Category    string `json:"category" validate:"required,min=1,max=15"`
	}

	AssignTicketByIdRequest struct {
		ID           uuid.UUID `json:"id" validate:"required"`
		UserID       int64     `json:"admin_id" validate:"required"`
		TechnicianID int64     `json:"technician_id" validate:"required"`
	}

	AssignTicketByIdResponse struct {
		Ticket TicketResponse `json:"ticket"`
	}

	CloseTicketByIdRequest struct {
		TicketId     uuid.UUID `json:"ticket_id" validate:"required"`
		UserID       int64     `json:"user_id" validate:"required"`
		TechnicianID int64     `json:"technician_id"`
		Activity     string    `json:"activity" validate:"required,min=1"`
	}
)
