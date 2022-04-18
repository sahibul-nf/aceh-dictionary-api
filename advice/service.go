package advice

import (
	"sort"

	jwd "github.com/jhvst/go-jaro-winkler-distance"
)

type Service interface {
	GetAdvices(query string) ([]Advice, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) GetAdvices(query string) ([]Advice, error) {

	advices := []Advice{}

	acehIndo, err := s.repository.FindLike(query)
	if err != nil {
		return advices, err
	}

	var result float64

	for _, v := range acehIndo {
		result = jwd.Calculate(query, v.Aceh)
		// fmt.Println(v.Aceh)
		// fmt.Println(result)
		// fmt.Println()

		if result >= 0.75 {
			advices = append(advices, Advice{
				ID:          v.ID,
				Aceh:        v.Aceh,
				Indonesia:   v.Indonesia,
				Similiarity: result,
			})
		}
	}

	sort.Slice(advices, func(i, j int) bool {
		return advices[i].Similiarity > advices[j].Similiarity
	})

	filters := []Advice{}

	if len(advices) > 5 {
		for i, v := range advices {
			if i < 5 {
				filters = append(filters, v)
			}
		}
		return filters, nil
	} else {
		return advices, nil
	}

}
