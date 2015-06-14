package mail

import (
	"time"

	"github.com/GeilMail/geilmail/storage/users"
	"github.com/jinzhu/gorm"
)

var db gorm.DB

type Mail struct {
	ID           uint
	IncomingDate time.Time
	Owner        users.User
	MailPath     string
	Unread       bool
}

func (m *Mail) GetContent() []byte {
	panic("GetContent")
	return []byte{}
}

func Prepare(gdb gorm.DB) error {
	db = gdb
	d := db.AutoMigrate(&Mail{})
	return d.Error
}
