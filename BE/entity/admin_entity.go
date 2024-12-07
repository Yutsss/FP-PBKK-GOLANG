package entity

type Admin struct {
	ID     uint `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID uint `gorm:"type:int;not null" json:"user_id"`
}
