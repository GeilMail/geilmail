package mail

import (
	"time"

	"gopkg.in/gorp.v1"
)

var db *gorp.DbMap

type Mail struct {
	ID           uint
	IncomingDate time.Time
	OwnerID      uint
	MailPath     string
	Unread       bool
}

func (m *Mail) GetContent() []byte {
	panic("GetContent")
	return []byte{}
}

func Prepare(dbm *gorp.DbMap) {
	db = dbm
	db.AddTableWithName(Mail{}, "mails")
}
