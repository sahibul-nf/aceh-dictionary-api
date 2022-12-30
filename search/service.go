package search

import (
	"sort"

	"github.com/agnivade/levenshtein"
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

			recommendationWords = append(recommendationWords, RecommendationWord{
				ID:          v.ID,
				Aceh:        v.Aceh,
				Indonesia:   v.Indonesia,
				English:     v.English,
				Similiarity: result,
			})
		}

		if algorithm == "lev" { // levenshtein distance
			lev := levenshtein.ComputeDistance(query, v.Aceh)
			// ubah ke persentase
			result = (float64(len(query)) - float64(lev)) / float64(len(query))

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

	if len(recommendationWords) > 10 {
		for i, v := range recommendationWords {
			if i < 10 {
				filters = append(filters, v)
			}
		}
		return filters, nil
	}

	return recommendationWords, nil
}
