package vocabulary

type Vocabulary struct {
	ID        int    `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`
	Aceh      string `gorm:"column:aceh;type:varchar;size:255" json:"aceh" binding:"required"`
	Indonesia string `gorm:"column:indonesia;type:varchar;size:255" json:"indonesia" binding:"required"`
	English   string `gorm:"column:english;type:varchar;size:255" json:"english" binding:"required"`
}
