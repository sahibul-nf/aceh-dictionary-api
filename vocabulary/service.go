package vocabulary

type Service interface {
	SaveData(input VocabularyInput) (bool, error)
	GetAllData() ([]Vocabulary, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) SaveData(input VocabularyInput) (bool, error) {

	var vocab Vocabulary
	vocab.Aceh = input.Aceh
	vocab.Indonesia = input.Indonesia

	isSave, err := s.repo.Save(vocab)
	if err != nil {
		return isSave, err
	}

	return isSave, nil
}

func (s *service) GetAllData() ([]Vocabulary, error) {

	data, err := s.repo.FindAll()
	if err != nil {
		return data, err
	}

	return data, nil
}
