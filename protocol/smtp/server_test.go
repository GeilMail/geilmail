package smtp

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
	"testing"
	"time"

	"github.com/GeilMail/geilmail/cfg"
	"github.com/GeilMail/geilmail/storage/mail"

	"github.com/facebookgo/ensure"
)

const smtpPort = 1587

func TestMain(m *testing.M) {
	Boot(&cfg.Config{
		SMTP: &cfg.SMTPConfig{
			ListenIP: "0.0.0.0",
			Port:     smtpPort,
		},
	})
	time.Sleep(time.Millisecond * 100) //TODO: replace that by a more sane mechanism
	ret := m.Run()
	ShutDown()
	os.Exit(ret)
}

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
	mailStorage = mail.GetInMemoryStorage()
	msgContent := []byte("This is a very specific message for the TestSMTPSubmission.\n")

	err := sendMail(fmt.Sprintf("localhost:%d", smtpPort), nil, "test@example.com", []string{"other@example.com"}, msgContent, false)
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
}

func TestStartTLS(t *testing.T) {
	msgContent := []byte("StartTLS test msg")
	err := sendMail(fmt.Sprintf("localhost:%d", smtpPort), nil, "test@example.com", []string{"test+subbox@example.com"}, msgContent, true)

	ensure.Nil(t, err)
}

// derived from golang's src/net/smtp/smtp.go, (http://golang.org/LICENSE)
func sendMail(addr string, a smtp.Auth, from string, to []string, msg []byte, enableTLS bool) error {
	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if enableTLS {
		if ok, _ := c.Extension("STARTTLS"); ok {
			config := &tls.Config{InsecureSkipVerify: true}
			if err = c.StartTLS(config); err != nil {
				return err
			}
		}
	}
	if a != nil {
		if err = c.Auth(a); err != nil {
			return err
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
