package mail

import (
	"sync"

	"github.com/GeilMail/geilmail/helpers"
	"github.com/GeilMail/geilmail/storage/users"
)

type InMemoryStorage struct {
	Mails    []*Mail
	MailsMtx sync.RWMutex
}

var ims *InMemoryStorage

func GetInMemoryStorage() *InMemoryStorage {
	if ims == nil {
		ims = &InMemoryStorage{Mails: []*Mail{}, MailsMtx: sync.RWMutex{}}
	}
	return ims
}

func (i *InMemoryStorage) MailDrop(content []byte, receiver helpers.MailAddress) error {
	panic("MailDrop")
	return nil
}

func (i *InMemoryStorage) GetUserMail(user users.User, path MailPath) ([]*Mail, error) {
	panic("GetUserMail")
	return nil, nil
}
