package pirate

type PirateService struct {
	repository PirateRepository
}

func NewService(repository PirateRepository) *PirateService {
	return &PirateService{repository: repository}
}

func (s *PirateService) AddPirate(name string, akumaNoMiID *uint, age int16, weapon []string, bounty float64, rank string) (*Pirate, error) {
	pirate := &Pirate{
		Name:        name,
		AkumaNoMiID: akumaNoMiID,
		Age:         age,
		Weapon:      weapon,
		Bounty:      bounty,
		Rank:        rank,
	}

	err := s.repository.Save(pirate)
	if err != nil {
		return nil, err
	}

	return pirate, nil
}

func (s *PirateService) FindByName(name string) (*Pirate, error) {
	return s.repository.FindByName(name)
}

func (s *PirateService) GetAllPirates() ([]Pirate, error) {
	return s.repository.GetAllPirates()
}

func (s *PirateService) FindById(id uint) (*Pirate, error) {
	return s.repository.FindById(id)
}
