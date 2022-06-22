package advice

import (
	"aceh-dictionary-api/vocabulary"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	FindLike(query string) ([]vocabulary.Vocabulary, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindLike(query string) ([]vocabulary.Vocabulary, error) {
	var vocabs []vocabulary.Vocabulary

	err := r.db.Where("aceh LIKE ?", "%"+strings.TrimSpace(query)+"%").Find(&vocabs).Error
	if err != nil {
		return vocabs, err
	}

	return vocabs, nil
}
