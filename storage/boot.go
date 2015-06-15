package storage

import (
	"database/sql"
	"log"
	"os"

	"github.com/GeilMail/geilmail/cfg"
	"github.com/GeilMail/geilmail/storage/mail"
	"github.com/GeilMail/geilmail/storage/users"
	"gopkg.in/gorp.v1"

	// driver import
	_ "github.com/mattn/go-sqlite3"
)

var dbMap *gorp.DbMap

func Boot(c *cfg.Config) {
	switch c.Storage.Provider {
	case "sqlite":
		db, err := sql.Open("sqlite3", c.Storage.SQLite.DBPath)
		if err != nil {
			panic(err)
		}
		dbMap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	default:
		panic("Invalid provider. Currently only 'sqlite' is supported.")
	}

	dbMap.TraceOn("[gorp]", log.New(os.Stdout, "geilmail: ", log.Lmicroseconds))

	users.Prepare(dbMap)
	mail.Prepare(dbMap)
	err := dbMap.CreateTablesIfNotExists()
	if err != nil {
		panic(err)
	}
}
