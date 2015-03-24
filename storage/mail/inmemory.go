package mail

import (
	"log"
	"sync"
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

func (i *InMemoryStorage) Store(m *Mail) error {
	i.MailsMtx.Lock()
	log.Printf("Storing %v\n", m)
	i.Mails = append(i.Mails, m)
	i.MailsMtx.Unlock()
	return nil
}
