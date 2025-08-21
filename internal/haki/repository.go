package haki

import "gorm.io/gorm"

type HakiRepository struct {
	db *gorm.DB
}

func NewHakiRepository(db *gorm.DB) *HakiRepository {
	return &HakiRepository{db: db}
}

func (r *HakiRepository) Save(haki *Haki) error {
	return r.db.Create(haki).Error
}

func (r *HakiRepository) GetAllHakis() ([]Haki, error) {
	var hakis []Haki
	if err := r.db.Find(&hakis).Error; err != nil {
		return nil, err
	}

	return hakis, nil
}
