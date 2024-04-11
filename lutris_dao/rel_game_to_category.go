package lutris_dao

type LutrisRelGameToCategory struct {
	GameId     int64 `gorm:"column:game_id;type:INTEGER;" json:"game_id"`
	CategoryId int64 `gorm:"column:category_id;type:INTEGER;" json:"category_id"`
}

func (l *LutrisRelGameToCategory) TableName() string {
	return "games_categories"
}
