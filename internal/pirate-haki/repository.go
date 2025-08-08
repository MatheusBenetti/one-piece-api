package piratehaki

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PirateHakiRepository struct {
	db *gorm.DB
}

func NewPirateHakiRepository(db *gorm.DB) *PirateHakiRepository {
	return &PirateHakiRepository{db: db}
}

func (r *PirateHakiRepository) AssignHakiToPirate(pirateID, hakiID uint, level string) error {
	ph := PirateHaki{
		PirateID: pirateID,
		HakiID:   hakiID,
		Level:    level,
	}

	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "pirate_id"}, {Name: "haki_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"level"}),
	}).Create(ph).Error
}

func (r *PirateHakiRepository) GetHakisByPirateID(pirateID uint) ([]PirateHaki, error) {
	var hakis []PirateHaki
	err := r.db.
		Preload("Haki").
		Where("pirate_id = ?", pirateID).
		Find(&hakis).Error

	return hakis, err
}

func (r *PirateHakiRepository) UpdateHakiLevel(pirateID, hakiID uint, newLevel string) error {
	return r.db.Model(&PirateHaki{}).
		Where("pirate_id = ? AND haki_id = ?", pirateID, hakiID).
		Update("level", newLevel).Error
}
