package advice

import (
	"aceh-dictionary-api/dictionary"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]dictionary.AcehIndo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]dictionary.AcehIndo, error) {
	var acehIndos []dictionary.AcehIndo

	err := r.db.Find(&acehIndos).Error
	if err != nil {
		return acehIndos, err
	}

	return acehIndos, nil
}
