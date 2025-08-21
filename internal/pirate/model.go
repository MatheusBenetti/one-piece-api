package pirate

import (
	akumanomi "github.com/MatheusBenetti/one-piece-api/internal/akuma-no-mi"
	haki "github.com/MatheusBenetti/one-piece-api/internal/haki"
	"gorm.io/gorm"
)

// TODO: Check the missing fields that can be used as a Pirate tipo Haki, battles, events, Crew

type Pirate struct {
	gorm.Model
	Name        string               `json:"name" gorm:"not null"`
	AkumaNoMiID *uint                `json:"akuma_no_mi_id"`
	AkumaNoMi   *akumanomi.AkumaNoMi `gorm:"foreignKey:AkumaNoMiID"`
	Haki        *haki.Haki           `gorm:"many2many=pirate_hakis;"`
	Age         int16                `json:"age"`
	Weapon      []string             `json:"weapon" gorm:"type:text[]"`
	Bounty      float64              `json:"bounty"`
	Rank        string               `json:"rank"`
}
