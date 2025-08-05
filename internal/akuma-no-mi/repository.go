package akumanomi

import "gorm.io/gorm"

type AkumaNoMiRepository struct {
	db *gorm.DB
}

func NewAkumaNoMiRepository(db *gorm.DB) *AkumaNoMiRepository {
	return &AkumaNoMiRepository{db: db}
}

func (r *AkumaNoMiRepository) Save(fruit *AkumaNoMi) error {
	return r.db.Create(fruit).Error
}

func (r *AkumaNoMiRepository) GetAllAkumaNoMi() ([]AkumaNoMi, error) {
	var fruits []AkumaNoMi
	if err := r.db.Find(&fruits).Error; err != nil {
		return nil, err
	}

	return fruits, nil
}

func (r *AkumaNoMiRepository) FindByName(name string) (*AkumaNoMi, error) {
	var fruit AkumaNoMi
	if err := r.db.Where("name = ?", name).First(&fruit).Error; err != nil {
		return nil, err
	}

	return &fruit, nil
}

func (r *AkumaNoMiRepository) FindById(id uint) (*AkumaNoMi, error) {
	var fruit AkumaNoMi
	if err := r.db.First(&fruit, id).Error; err != nil {
		return nil, err
	}

	return &fruit, nil
}
