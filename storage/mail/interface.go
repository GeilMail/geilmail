package mail

import (
	"github.com/GeilMail/geilmail/helpers"
	"github.com/GeilMail/geilmail/storage/users"
)

type Storage interface {
	MailDrop(content []byte, receiver helpers.MailAddress) error
	GetUserMail(user users.User, path MailPath) ([]*Mail, error)
}
