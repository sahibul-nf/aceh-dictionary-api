package bookmark

import "gorm.io/gorm"

type Repository interface {
	Save(bookmark Bookmark) (Bookmark, error)
	FindByID(ID int) (Bookmark, error)
	FindByUserID(userID int) ([]Bookmark, error)
	Delete(bookmark Bookmark) (Bookmark, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(bookmark Bookmark) (Bookmark, error) {

	err := r.db.Create(&bookmark).Error
	if err != nil {
		return bookmark, err
	}

	return bookmark, nil
}

func (r *repository) FindByID(ID int) (Bookmark, error) {
	var bookmark Bookmark

	err := r.db.Where("id = ?", ID).Find(&bookmark).Error
	if err != nil {
		return bookmark, err
	}

	return bookmark, nil
}

func (r *repository) FindByUserID(userID int) ([]Bookmark, error) {
	var bookmarks []Bookmark

	err := r.db.Where("user_id = ?", userID).Find(&bookmarks).Error
	if err != nil {
		return bookmarks, err
	}

	return bookmarks, nil
}

func (r *repository) Delete(bookmark Bookmark) (Bookmark, error) {
	err := r.db.Delete(&bookmark).Error
	if err != nil {
		return bookmark, err
	}

	return bookmark, nil
}
