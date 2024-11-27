package command

import (
	"github.com/Yutsss/FP-PBKK-GOLANG/BE/migrations"
	"gorm.io/gorm"
	"log"
	"os"
)

func Commands(db *gorm.DB) bool {
	migrate := false
	run := true

	for _, arg := range os.Args[1:] {
		if arg == "--migrate" {
			migrate = true
		}
	}

	if migrate {
		if err := migrations.Migrate(db); err != nil {
			log.Fatalf("error migration: %v", err)
		}
		log.Println("migration completed successfully")
		return false
	}

	if run {
		return true
	}

	return false
}
