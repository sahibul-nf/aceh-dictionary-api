package advice

import (
	"aceh-dictionary-api/dictionary"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	FindLike(query string) ([]dictionary.AcehIndo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindLike(query string) ([]dictionary.AcehIndo, error) {
	var acehIndos []dictionary.AcehIndo

	err := r.db.Where("aceh LIKE ?", "%" + strings.TrimSpace(query) + "%").Find(&acehIndos).Error
	if err != nil {
		return acehIndos, err
	}
	
	return acehIndos, nil
}