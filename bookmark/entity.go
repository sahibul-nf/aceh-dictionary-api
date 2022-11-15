package bookmark

import "time"

type Bookmark struct {
	ID           int       `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`
	UserID       int       `gorm:"column:user_id;type:int" json:"user_id"`
	DictionaryID int       `gorm:"column:dictionary_id;type:int" json:"dictionary_id"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}
