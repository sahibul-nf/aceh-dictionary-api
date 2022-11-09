package dictionary

type Dictionary struct {
	ID        int    `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`
	Aceh      string `gorm:"column:aceh;size:255" json:"aceh" binding:"required"`
	Indonesia string `gorm:"column:indonesia;size:255" json:"indonesia" binding:"required"`
	English   string `gorm:"column:english;size:255" json:"english" binding:"required"`
	CreatedAt string `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}
