package user

import "time"

type User struct {
	ID           int       `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name         string    `gorm:"column:name;size:255"`
	Email        string    `gorm:"uniqueIndex:email;size:255"`
	PasswordHash string    `gorm:"column:password_hash;size:255"`
	AvatarURL    string    `gorm:"column:avatar_url;size:255"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp"`
}
