package entity

import (
	"github.com/google/uuid"
)

type Ticket struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID       uint      `gorm:"type:int;not null" json:"user_id"`
	AdminID      uint      `gorm:"type:int" json:"admin_id"`
	LogID        uint      `gorm:"type:int" json:"log_id"`
	TechnicianID uint      `gorm:"type:int" json:"technician_id"`
	Title        string    `gorm:"type:varchar(255);not null" json:"title"`
	Description  string    `gorm:"type:text;not null" json:"description"`
	Category     string    `gorm:"type:varchar(15);not null" json:"category"`
	Status       string    `gorm:"type:varchar(15);not null;default: 'open'" json:"status"`
}
