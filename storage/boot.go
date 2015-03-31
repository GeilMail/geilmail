package storage

import (
	"github.com/GeilMail/geilmail/configuration"
)

func Boot(c *configuration.Config) {
	openSQLiteConnection(c.SQLite)
}
