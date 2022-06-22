package vocabulary

import "gorm.io/gorm"

type Repository interface {
	Save(vocab Vocabulary) (bool, error)
	FindAll() ([]Vocabulary, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(vocab Vocabulary) (bool, error) {

	err := r.db.Create(&vocab).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository) FindAll() ([]Vocabulary, error) {
	var vocabs []Vocabulary

	err := r.db.Find(&vocabs).Error
	if err != nil {
		return vocabs, err
	}

	return vocabs, nil
}
