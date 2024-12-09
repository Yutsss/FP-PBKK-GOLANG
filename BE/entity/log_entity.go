package entity

type Log struct {
	ID           int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	TechnicianID int64  `gorm:"type:int;not null" json:"technician_id"`
	Activity     string `gorm:"type:text;not null" json:"activity"`
	Ticket       Ticket `gorm:"foreignKey:LogID;references:ID" json:"ticket"`
}
