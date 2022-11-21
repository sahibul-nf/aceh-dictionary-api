package dictionary

import "time"

type Dictionary struct {
	ID        int       `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`
	Aceh      string    `gorm:"column:aceh;type:varchar;size:255" json:"aceh"`
	Indonesia string    `gorm:"column:indonesia;type:varchar;size:255" json:"indonesia"`
	English   string    `gorm:"column:english;type:varchar;size:255" json:"english"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
