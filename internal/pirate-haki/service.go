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

func (s *PirateHakiService) GetHakisByPirateID(pirateID uint) ([]PirateHakiResponse, error) {
	links, err := s.repository.GetHakisByPirateID(pirateID)
	if err != nil {
		return nil, err
	}

	out := make([]PirateHakiResponse, 0, len(links))
	for _, ph := range links {
		out = append(out, PirateHakiResponse{
			HakiID:   ph.HakiID,
			HakiName: ph.Haki.Name,
			Level:    ph.Level,
		})
	}

	return out, nil
}
