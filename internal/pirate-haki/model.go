package piratehaki

import (
	haki "github.com/MatheusBenetti/one-piece-api/internal/haki"
	pirate "github.com/MatheusBenetti/one-piece-api/internal/pirate"
)

type PirateHaki struct {
	PirateID uint   `gorm:"primaryKey"`
	HakiID   uint   `gorm:"primaryKey"`
	Level    string `gorm:"type:varchar(20);not null"`

	Pirate pirate.Pirate `gorm:"foreignKey=PirateID;constraint:OnDelete:CASCADE"`
	Haki   haki.Haki     `gorm:"foreignKey=HakiID;constraint:OnDelete:CASCADE"`
}
