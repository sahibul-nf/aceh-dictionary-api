package search

import (
	"aceh-dictionary-api/dictionary"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	FindLike(query string) ([]dictionary.Dictionary, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindLike(query string) ([]dictionary.Dictionary, error) {
	var dictionary []dictionary.Dictionary

	err := r.db.Where("aceh LIKE ?", "%"+strings.TrimSpace(query)+"%").Find(&dictionary).Error
	if err != nil {
		return dictionary, err
	}

	return dictionary, nil
}
