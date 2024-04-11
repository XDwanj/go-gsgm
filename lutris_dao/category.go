package lutris_dao

type LutrisCategory struct {
	Id          int64        `gorm:"column:id;type:INTEGER;primaryKey;" json:"id"`
	Name        string       `gorm:"column:name;type:TEXT;" json:"name"`
}

func (l *LutrisCategory) TableName() string {
	return "categories"
}
