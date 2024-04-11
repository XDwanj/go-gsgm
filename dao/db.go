package dao

import (
	"database/sql"
	"strings"

	"github.com/XDwanj/go-gsgm/config"
	"github.com/XDwanj/go-gsgm/logger"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qustavo/sqlhooks/v2"
)

var LuDb *sqlx.DB

func init() {
	// wrap
	sql.Register("sqlite3_with_hooks", sqlhooks.Wrap(&sqlite3.SQLiteDriver{}, &Hooks{}))

	db, err := sqlx.Connect("sqlite3_with_hooks", config.PgaDbPath)
	if err != nil {
		logger.Erro(err)
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)

	LuDb = db
}
