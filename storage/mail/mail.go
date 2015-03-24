package mail

import (
	"time"
)

type Mail struct {
	IncomingDate time.Time
	Recipient    string
	Sender       string
	Content      []byte
}
