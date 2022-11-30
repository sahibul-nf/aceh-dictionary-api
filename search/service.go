package search

import (
	"sort"

	"github.com/agext/levenshtein"
	jwd "github.com/jhvst/go-jaro-winkler-distance"
)

type Service interface {
	GetRecommendationWords(query string, algorithm string) ([]RecommendationWord, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) GetRecommendationWords(query string, algorithm string) ([]RecommendationWord, error) {

	recommendationWords := []RecommendationWord{}

	dictionaries, err := s.repository.FindAll()
	if err != nil {
		return recommendationWords, err
	}

	for _, v := range dictionaries {
		var result float64

		if algorithm == "jwd" { // jaro winkler distance
			result = jwd.Calculate(query, v.Aceh)
		}

		if algorithm == "lev" { // levenshtein distance
			result = levenshtein.Match(query, v.Aceh, nil)
		}

		if result >= 0.75 {
			recommendationWords = append(recommendationWords, RecommendationWord{
				ID:          v.ID,
				Aceh:        v.Aceh,
				Indonesia:   v.Indonesia,
				English:     v.English,
				Similiarity: result,
			})
		}
	}

	sort.Slice(recommendationWords, func(i, j int) bool {
		return recommendationWords[i].Similiarity > recommendationWords[j].Similiarity
	})

	filters := []RecommendationWord{}

	if len(recommendationWords) > 5 {
		for i, v := range recommendationWords {
			if i < 5 {
				filters = append(filters, v)
			}
		}
		return filters, nil
	} else {
		return recommendationWords, nil
	}

}
