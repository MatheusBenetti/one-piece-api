package config

import (
	akumaNoMi "github.com/MatheusBenetti/one-piece-api/internal/akuma-no-mi"
	pirate "github.com/MatheusBenetti/one-piece-api/internal/pirate"
	user "github.com/MatheusBenetti/one-piece-api/internal/user"
	"gorm.io/gorm"
)

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&akumaNoMi.AkumaNoMi{}, &pirate.Pirate{}, &user.User{})
}
