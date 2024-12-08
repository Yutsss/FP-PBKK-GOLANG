package entity

type Admin struct {
	ID      int64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID  int64    `gorm:"type:int;not null" json:"user_id"`
	Tickets []Ticket `gorm:"foreignKey:AdminID;references:ID" json:"tickets"`
}
