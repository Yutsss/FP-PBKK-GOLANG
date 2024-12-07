package entity

type Technician struct {
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         uint   `gorm:"type:int;not null" json:"user_id"`
	Specialization string `gorm:"type:varchar(255);not null" json:"specialization"`
}
