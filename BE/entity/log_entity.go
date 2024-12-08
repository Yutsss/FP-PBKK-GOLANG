package entity

type Log struct {
	ID           int64  `gorm:"primaryKey" json:"id"`
	TicketID     int64  `gorm:"type:int;not null" json:"ticket_id"`
	TechnicianID int64  `gorm:"type:int;not null" json:"technician_id"`
	Activity     string `gorm:"type:text;not null" json:"activity"`
}
