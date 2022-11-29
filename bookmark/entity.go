package bookmark

import (
	"aceh-dictionary-api/dictionary"
	"aceh-dictionary-api/user"
	"time"
)

type Bookmark struct {
	ID           int       `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	UserID       int       `gorm:"column:user_id;type:int" `
	DictionaryID int       `gorm:"column:dictionary_id;type:int"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp" `
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp"`
	User         user.User
	Dictionary   dictionary.Dictionary
}
