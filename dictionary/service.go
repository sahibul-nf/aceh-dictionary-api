package dictionary

type Service interface {
	SaveData(input AcehIndo) (bool, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) SaveData(input AcehIndo) (bool, error) {

	isSave, err := s.repo.Save(input)
	if err != nil {
		return isSave, err
	}

	return isSave, nil
}
