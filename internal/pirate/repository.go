package pirate

import "gorm.io/gorm"

type PirateRepository struct {
	db *gorm.DB
}

func NewPirateRepository(db *gorm.DB) *PirateRepository {
	return &PirateRepository{db: db}
}

func (r *PirateRepository) Save(pirate *Pirate) error {
	return r.db.Create(pirate).Error
}

func (r *PirateRepository) GetAllPirates() ([]Pirate, error) {
	var pirates []Pirate
	if err := r.db.Find(&pirates).Error; err != nil {
		return nil, err
	}

	return pirates, nil
}

func (r *PirateRepository) FindByName(name string) (*Pirate, error) {
	var pirate Pirate
	if err := r.db.Where("name = ?", name).First(&pirate).Error; err != nil {
		return nil, err
	}

	return &pirate, nil
}

func (r *PirateRepository) FindById(id uint) (*Pirate, error) {
	var pirate Pirate
	if err := r.db.First(&pirate, id).Error; err != nil {
		return nil, err
	}

	return &pirate, nil
}
