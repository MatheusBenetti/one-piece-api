package haki

type HakiService struct {
	repository *HakiRepository
}

func NewHakiService(repository *HakiRepository) *HakiService {
	return &HakiService{repository: repository}
}

func (s *HakiService) CreateHaki(name string) (*Haki, error) {
	haki := &Haki{Name: name}
	if err := s.repository.Save(haki); err != nil {
		return nil, err
	}

	return haki, nil
}

func (s *HakiService) GetAllHakis() ([]Haki, error) {
	return s.repository.GetAllHakis()
}
