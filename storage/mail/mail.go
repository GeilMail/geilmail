package mail

import (
	"time"

	"github.com/GeilMail/geilmail/storage/users"
	"github.com/landjur/go-uuid"
)

type MailPath string
type MailID uuid.UUID

type Mail struct {
	ID           MailID
	IncomingDate time.Time
	Owner        users.User
	MailPath     MailPath
	Unread       bool
}

func (m *Mail) GetContent() []byte {
	panic("GetContent")
	return []byte{}
}
