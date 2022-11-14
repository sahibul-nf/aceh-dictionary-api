package dictionary

type Service interface {
	SaveWord(input DictionaryInput) (Dictionary, error)
	GetWords() ([]Dictionary, error)
	GetWord(id int) (Dictionary, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) SaveWord(input DictionaryInput) (Dictionary, error) {

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

func (s *service) GetWords() ([]Dictionary, error) {

	dictionaries, err := s.repo.FindAll()
	if err != nil {
		return dictionaries, err
	}

	return dictionaries, nil
}

func (s *service) GetWord(id int) (Dictionary, error) {

	dictionary, err := s.repo.FindByID(id)
	if err != nil {
		return dictionary, err
	}

	return dictionary, nil
}
