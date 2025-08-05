package akumanomi

type AkumaNoMiService struct {
	repository AkumaNoMiRepository
}

func newService(repository AkumaNoMiRepository) *AkumaNoMiService {
	return &AkumaNoMiService{repository: repository}
}

func (s *AkumaNoMiService) AddAkumaNoMi(name, model, meaning, description, sort string) (*AkumaNoMi, error) {
	akumaNoMi := &AkumaNoMi{
		Name:        name,
		Model:       model,
		Meaning:     meaning,
		Description: description,
		Sort:        sort,
	}

	if err := s.repository.Save(akumaNoMi); err != nil {
		return nil, err
	}

	return akumaNoMi, nil
}

func (s *AkumaNoMiService) FindByName(name string) (*AkumaNoMi, error) {
	return s.repository.FindByName(name)
}

func (s *AkumaNoMiService) GetAllAkumaNoMi() ([]AkumaNoMi, error) {
	return s.repository.GetAllAkumaNoMi()
}

func (s *AkumaNoMiService) FindById(id uint) (*AkumaNoMi, error) {
	return s.repository.FindById(id)
}
