package storage

import (
	"github.com/GeilMail/geilmail/cfg"
)

func Boot(c *cfg.Config) {
	openSQLiteConnection(c.SQLite)
}
