package dictionary

import "gorm.io/gorm"

type Repository interface {
	Save(acehIndos AcehIndo) (bool, error)
	FindAll() ([]AcehIndo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(acehIndos AcehIndo) (bool, error) {

	err := r.db.Create(&acehIndos).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository) FindAll() ([]AcehIndo, error) {
	var acehIndos []AcehIndo
	
	err := r.db.Find(&acehIndos).Error
	if err != nil {
		return acehIndos, err
	}
	
	return acehIndos, nil
}
