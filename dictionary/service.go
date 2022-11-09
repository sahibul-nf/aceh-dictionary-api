package dictionary

type Service interface {
	SaveData(input DictionaryInput) (Dictionary, error)
	GetAllData() ([]Dictionary, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) SaveData(input DictionaryInput) (Dictionary, error) {

	var dictionary Dictionary
	dictionary.Aceh = input.Aceh
	dictionary.Indonesia = input.Indonesia
	dictionary.English = input.English

	newDictionary, err := s.repo.Save(dictionary)
	if err != nil {
		return newDictionary, err
	}

	return newDictionary, nil
}

func (s *service) GetAllData() ([]Dictionary, error) {

	dictionaries, err := s.repo.FindAll()
	if err != nil {
		return dictionaries, err
	}

	return dictionaries, nil
}
