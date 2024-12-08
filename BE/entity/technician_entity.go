package entity

type Technician struct {
	ID             int64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         int64    `gorm:"type:int;not null" json:"user_id"`
	Specialization string   `gorm:"type:varchar(255);not null" json:"specialization"`
	Tickets        []Ticket `gorm:"foreignKey:TechnicianID;references:ID" json:"tickets"`
}
