package config

import (
	akumaNoMi "github.com/MatheusBenetti/one-piece-api/internal/akuma-no-mi"
	haki "github.com/MatheusBenetti/one-piece-api/internal/haki"
	pirate "github.com/MatheusBenetti/one-piece-api/internal/pirate"
	piratehaki "github.com/MatheusBenetti/one-piece-api/internal/pirate-haki"
	user "github.com/MatheusBenetti/one-piece-api/internal/user"
	"gorm.io/gorm"
)

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&akumaNoMi.AkumaNoMi{}, &haki.Haki{}, &pirate.Pirate{}, &piratehaki.PirateHaki{}, &user.User{})
}
