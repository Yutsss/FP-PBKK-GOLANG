package migrations

import (
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.User{})

	if err != nil {
		return err
	}

	return nil
}
