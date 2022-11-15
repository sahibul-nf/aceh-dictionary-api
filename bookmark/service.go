package bookmark

type Service interface {
	MarkWord(input MarkWordInput) (Bookmark, error)
	UnmarkWord(input MarkWordInput) error
	FindByUserIDAndDictionaryID(userID int, dictionaryID int) (Bookmark, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) MarkWord(input MarkWordInput) (Bookmark, error) {
	bookmarkInput := Bookmark{
		UserID:       input.User.ID,
		DictionaryID: input.DictionaryID,
	}

	newBookmark, err := s.repository.Save(bookmarkInput)
	if err != nil {
		return newBookmark, err
	}

	return newBookmark, nil
}

func (s *service) UnmarkWord(input MarkWordInput) error {
	err := s.repository.DeleteByUserIDAndDictionaryID(input.User.ID, input.DictionaryID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) FindByUserIDAndDictionaryID(userID int, dictionaryID int) (Bookmark, error) {
	markedWord, err := s.repository.FindByUserIDAndDictionaryID(userID, dictionaryID)
	if err != nil {
		return markedWord, err
	}

	return markedWord, nil
}
