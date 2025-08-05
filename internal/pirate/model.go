package pirate

import akumanomi "github.com/MatheusBenetti/one-piece-api/internal/akuma-no-mi"

type Pirate struct {
	ID          uint                 `json:"id" gorm:"primaryKey"`
	Name        string               `json:"name" gorm:"not null"`
	AkumaNoMiID *uint                `json:"akuma_no_mi_id"`
	AkumaNoMi   *akumanomi.AkumaNoMi `gorm:"foreignKey:AkumaNoMiID"`
}
