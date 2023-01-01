package search

type RecommendationWord struct {
	ID          int     `json:"id"`
	Aceh        string  `json:"aceh"`
	Indonesia   string  `json:"indonesia"`
	English     string  `json:"english"`
	Similiarity float64 `json:"similiarity"`
}