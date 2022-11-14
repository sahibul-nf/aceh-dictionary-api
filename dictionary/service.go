package dictionary

import (
	"aceh-dictionary-api/unsplash"
	"strings"
)

type Service interface {
	SaveWord(input DictionaryInput) (Dictionary, error)
	GetWords() (DictionariesFormatter, error)
	GetWord(id int) (DictionaryFormatter, error)
}

type service struct {
	repo         Repository
	unsplashRepo unsplash.Repository
}

func NewService(r Repository, unsplashRepo unsplash.Repository) *service {
	return &service{r, unsplashRepo}
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

func (s *service) GetWords() (DictionariesFormatter, error) {
	var dictionariesFormatter DictionariesFormatter
	dictionaries, err := s.repo.FindAll()
	if err != nil {
		return dictionariesFormatter, err
	}

	dictionariesFormatter = FormatDictionariesWithTotal(dictionaries)

	// ! handle unsplash image
	// for _, word := range dictionariesFormatter.Words {
	// 	wordSearchImage := strings.Split(word.English, ",")[0]
	// 	image, err := s.unsplashRepo.GetPhotoByKeyword(wordSearchImage, 1, "squarish")
	// 	if err != nil {
	// 		if err.Error() == "keyword not found" {
	// 			continue
	// 		}
	// 		return dictionariesFormatter, err
	// 	}

	// 	if len(image) > 0 {
	// 		word.ImagesURL = append(word.ImagesURL, image[0].Urls.Regular)
	// 	}
	// }

	return dictionariesFormatter, nil
}

func (s *service) GetWord(id int) (DictionaryFormatter, error) {
	var dictionaryFormatter DictionaryFormatter

	dictionary, err := s.repo.FindByID(id)
	if err != nil {
		return dictionaryFormatter, err
	}

	dictionaryFormatter = FormatDictionary(dictionary)

	wordSearchImage := strings.Split(dictionary.English, ",")[0]
	image, err := s.unsplashRepo.GetPhotoByKeyword(wordSearchImage, 3, "portrait")
	if err != nil {
		return dictionaryFormatter, err
	}

	if len(image) > 0 {
		for _, i := range image {
			dictionaryFormatter.ImagesURL = append(dictionaryFormatter.ImagesURL, i.Urls.Regular)
		}
	}

	return dictionaryFormatter, nil
}
