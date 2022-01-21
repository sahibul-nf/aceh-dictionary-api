package dictionary

import "gorm.io/gorm"

type Repository interface {
	Save(acehIndos []AcehIndo) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(acehIndos []AcehIndo) (bool, error) {

	err := r.db.Create(&acehIndos).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
