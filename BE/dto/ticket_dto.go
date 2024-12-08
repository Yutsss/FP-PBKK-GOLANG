package dto

import (
	"github.com/google/uuid"
)

type (
	TicketResponse struct {
		ID           uuid.UUID `json:"id"`
		UserID       uint      `json:"user_id"`
		AdminID      uint      `json:"admin_id"`
		LogID        uint      `json:"log_id"`
		TechnicianID uint      `json:"technician_id"`
		Title        string    `json:"title"`
		Description  string    `json:"description"`
		Category     string    `json:"category"`
		Status       string    `json:"status"`
	}

	CreateTicketRequest struct {
		UserID      uint   `json:"user_id" validate:"required"`
		Title       string `json:"title" validate:"required,min=1,max=255"`
		Description string `json:"description" validate:"required,min=1"`
		Category    string `json:"category" validate:"required,min=1,max=15"`
	}

	CreateTicketResponse struct {
		ID          uuid.UUID `json:"id"`
		UserID      uint      `json:"user_id"`
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
		Tickets []TicketResponse `json:"tickets"`
	}

	GetTicketByUserIDRequest struct {
		UserID uint `json:"user_id" validate:"required"`
	}

	GetTicketByUserIDResponse struct {
		Tickets []TicketResponse `json:"tickets"`
	}
)
