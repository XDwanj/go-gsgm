package lutris_dao

/*
CREATE TABLE games_categories (

	game_id INTEGER,
	category_id INTEGER

)
*/
type LutrisRelGameToCategory struct {
	GameId     int64 `json:"game_id"`
	CategoryId int64 `json:"category_id"`
}
