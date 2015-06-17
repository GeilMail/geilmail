package smtp

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
	"testing"

	"github.com/GeilMail/geilmail/cfg"
	"github.com/GeilMail/geilmail/storage"

	"github.com/facebookgo/ensure"
)

const testSMTPPort = 1587

var testConfStruct = &cfg.Config{
	SMTP: cfg.SMTPConfig{
		ListenIP: "0.0.0.0",
		Port:     testSMTPPort,
	},
	TLS: cfg.TLSConfig{},
	Storage: cfg.StorageConfig{
		Provider: "sqlite",
		SQLite: struct{ DBPath string }{
			DBPath: "",
		},
	},
}

func TestMain(m *testing.M) {
	testDBPath := "test.db"
	conf := testConfStruct
	conf.Storage.SQLite.DBPath = testDBPath
	storage.Boot(conf)
	rdy := Boot(conf)
	<-rdy
	ret := m.Run()
	ShutDown()
	os.Remove(testDBPath)
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
	msgContent := []byte("This is a very specific message for the TestSMTPSubmission.\n")

	err := sendMail(fmt.Sprintf("localhost:%d", testSMTPPort), nil, "test@example.com", []string{"other@example.com"}, msgContent, false)
	ensure.Nil(t, err)

	//TODO: test if it has been delivered
}

func TestStartTLS(t *testing.T) {
	msgContent := []byte("StartTLS test msg")
	err := sendMail(fmt.Sprintf("localhost:%d", testSMTPPort), nil, "test@example.com", []string{"test+subbox@example.com"}, msgContent, true)

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
