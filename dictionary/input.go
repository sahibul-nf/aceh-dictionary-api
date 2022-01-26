package dictionary

type DictionaryInput struct {
	Aceh      string `json:"aceh" binding:"required"`
	Indonesia string `json:"indonesia" binding:"required"`
}
