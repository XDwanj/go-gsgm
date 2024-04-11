package lutris_dao

/*
CREATE TABLE categories (

	id INTEGER PRIMARY KEY,
	name TEXT UNIQUE

)
*/
type LutrisCategory struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
