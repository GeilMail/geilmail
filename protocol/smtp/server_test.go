package smtp

import (
	"fmt"
	"net/smtp"
	"testing"

	"github.com/GeilMail/geilmail/storage/mail"

	"github.com/facebookgo/ensure"
)

func eq(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestSMTPSubmission(t *testing.T) {
	Boot()
	mailStorage = mail.GetInMemoryStorage()

	msgContent := []byte("This is a very specific message for the TestSMTPSubmission.\n")
	err := smtp.SendMail(fmt.Sprintf("localhost:%d", smtpPort), nil, "test@example.com", []string{"other@example.com"}, msgContent)
	ensure.Nil(t, err)

	found := false
	mis := mailStorage.(*mail.InMemoryStorage)
	mis.MailsMtx.RLock() // just to be sure
	for _, msg := range mis.Mails {
		if eq(msg.Content, msgContent) {
			found = true
			break
		}
	}
	mis.MailsMtx.RUnlock()
	ensure.True(t, found)

	ShutDown()
}
