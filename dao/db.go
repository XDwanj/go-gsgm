package dao

import (
	"github.com/XDwanj/go-gsgm/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	// "gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
)

var LutrisDb *gorm.DB

func init() {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(
		sqlite.Open(config.PgaDbPath),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)
	if err != nil {
		panic("lutris 数据库未初始化")
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(1)
	sqlDb.SetMaxIdleConns(1)
	LutrisDb = db
}
