package entity

import (
	"database/sql"
	"github.com/google/uuid"
)

type Ticket struct {
	ID           uuid.UUID     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID       int64         `gorm:"type:int;not null;index" json:"user_id"`
	AdminID      sql.NullInt64 `gorm:"type:int;default null" json:"admin_id"`
	LogID        sql.NullInt64 `gorm:"type:int;default null" json:"log_id"`
	TechnicianID sql.NullInt64 `gorm:"type:int;default null" json:"technician_id"`
	Title        string        `gorm:"type:varchar(255);not null" json:"title"`
	Description  string        `gorm:"type:text;not null" json:"description"`
	Category     string        `gorm:"type:varchar(15);not null" json:"category"`
	Status       string        `gorm:"type:varchar(15);not null;default: 'open'" json:"status"`
}
