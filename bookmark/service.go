package bookmark

type Service interface {
	MarkWord(input MarkWordInput) (Bookmark, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) MarkWord(input MarkWordInput) (Bookmark, error) {
	bookmark := Bookmark{
		UserID:       input.User.ID,
		DictionaryID: input.DictionaryID,
	}

	bookmark, err := s.repository.Save(bookmark)
	if err != nil {
		return bookmark, err
	}

	return bookmark, nil
}
