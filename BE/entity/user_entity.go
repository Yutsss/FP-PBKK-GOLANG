package entity

import (
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/utility"
	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string `gorm:"type:varchar(255);not null" json:"name"`
	CompleteName string `gorm:"type:varchar(255);not null" json:"complate_name"`
	Email        string `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	Password     string `gorm:"type:varchar(255);not null" json:"password"`
	PhoneNumber  string `gorm:"type:varchar(255);not null" json:"phone_number"`
	Address      string `gorm:"type:varchar(255);not null" json:"address"`
	Role         string `gorm:"type:varchar(10);not null;default:'user'" json:"role"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var err error
	u.Password, err = utility.HashPassword(u.Password)
	if err != nil {
		return err
	}

	return nil
}
