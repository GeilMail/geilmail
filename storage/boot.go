package storage

import (
	"github.com/GeilMail/geilmail/cfg"
	"github.com/GeilMail/geilmail/storage/users"

	"github.com/jinzhu/gorm"
	// driver import
	_ "github.com/mattn/go-sqlite3"
)

var db gorm.DB

func Boot(c *cfg.Config) {
	var err error

	switch c.Storage.Provider {
	case "sqlite":
		db, err = gorm.Open("sqlite3", c.Storage.SQLite.DBPath)
		if err != nil {
			panic(err)
		}
	default:
		panic("Invalid provider. Currently only 'sqlite' is supported.")
	}

	users.Prepare(db)
}
