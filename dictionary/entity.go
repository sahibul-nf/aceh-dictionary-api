package dictionary

type AcehIndo struct {
	ID        int    `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Aceh      string `gorm:"column:aceh;type:varchar;size:255"`
	Indonesia string `gorm:"column:indonesia;type:varchar;size:255"`
}
