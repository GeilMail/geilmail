package storage

import (
	"database/sql"

	"github.com/GeilMail/geilmail/cfg"

	_ "github.com/mattn/go-sqlite3"
)

var SQLiteConn *sql.DB

func openSQLiteConnection(sqlconf cfg.SQLiteConfig) {
	var err error
	SQLiteConn, err = sql.Open("sqlite3", sqlconf.DBPath)
	if err != nil {
		panic(err)
	}
}
