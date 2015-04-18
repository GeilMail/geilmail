package mail

import (
	"sync"
	"time"

	"github.com/GeilMail/geilmail/helpers"
	"github.com/GeilMail/geilmail/storage/users"

	"github.com/landjur/go-uuid"
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
	mailID, err := uuid.NewTimeBased()
	if err != nil {
		return err
	}
	m := Mail{
		ID:           MailID(mailID),
		IncomingDate: time.Now(),
		MailPath:     MailPath("/"),
		Unread:       true,
	} //TODO: owner, path etc

	i.MailsMtx.Lock()
	i.Mails = append(i.Mails, &m)
	i.MailsMtx.Unlock()
	return nil
}

func (i *InMemoryStorage) GetUserMail(user users.User, path MailPath) ([]*Mail, error) {
	panic("GetUserMail")
	return nil, nil
}
