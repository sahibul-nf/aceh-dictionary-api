package dictionary

import "gorm.io/gorm"

type Repository interface {
	Save(dictionary Dictionary) (Dictionary, error)
	FindAll() ([]Dictionary, error)
	FindByID(id int) (Dictionary, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(dictionary Dictionary) (Dictionary, error) {

	err := r.db.Create(&dictionary).Error
	if err != nil {
		return dictionary, err
	}

	return dictionary, nil
}

func (r *repository) FindAll() ([]Dictionary, error) {
	var dictionaries []Dictionary

	err := r.db.Find(&dictionaries).Error
	if err != nil {
		return dictionaries, err
	}

	return dictionaries, nil
}

func (r *repository) FindByID(id int) (Dictionary, error) {
	var dictionary Dictionary

	err := r.db.Where("id = ?", id).Find(&dictionary).Error
	if err != nil {
		return dictionary, err
	}

	return dictionary, nil
}
