package entity

type Log struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	TicketID     uint   `gorm:"type:int;not null" json:"ticket_id"`
	TechnicianID uint   `gorm:"type:int;not null" json:"technician_id"`
	Activity     string `gorm:"type:text;not null" json:"activity"`
}
