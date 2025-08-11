package piratehaki

type PirateHakiService struct {
	repository *PirateHakiRepository
}

func NewPirateHakiService(repository *PirateHakiRepository) *PirateHakiService {
	return &PirateHakiService{repository: repository}
}

func (s *PirateHakiService) AssignHakiToPirate(pirateID, hakiID uint, level string) error {
	return s.repository.AssignHakiToPirate(pirateID, hakiID, level)
}

func (s *PirateHakiService) UpdateHakiLevel(pirateID, hakiID uint, newLevel string) error {
	return s.repository.UpdateHakiLevel(pirateID, hakiID, newLevel)
}
