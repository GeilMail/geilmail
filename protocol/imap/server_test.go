package imap

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/GeilMail/geilmail/cfg"
	"github.com/GeilMail/geilmail/helpers"
	"github.com/GeilMail/geilmail/storage"
	"github.com/GeilMail/geilmail/storage/users"
	"github.com/facebookgo/ensure"
	"github.com/mxk/go-imap/imap"
)

const imapTestPort = 1143

func TestMain(m *testing.M) {
	testDBPath := "test.db"
	conf := &cfg.Config{
		IMAP: cfg.IMAPConfig{
			ListenIP: "0.0.0.0",
			Port:     imapTestPort,
		},
		TLS: cfg.TLSConfig{},
		Storage: cfg.StorageConfig{
			Provider: "sqlite",
			SQLite: struct{ DBPath string }{
				DBPath: testDBPath,
			},
		},
	}
	rdy := Boot(conf)
	storage.Boot(conf)
	err := users.New(helpers.MailAddress("test@example.com"), "1234")
	log.Println(">>>>>>>>>>")
	log.Println(users.CheckPassword("test@example.com", []byte("1234")))
	if err != nil {
		panic(err)
	}
	<-rdy
	ret := m.Run()
	ShutDown()
	os.Remove(testDBPath)
	os.Exit(ret)
}

// This is an integration test for the IMAP workflow.
func TestIMAPFlow(t *testing.T) {
	c, err := imap.Dial(fmt.Sprintf("localhost:%d", imapTestPort))
	ensure.Nil(t, err)
	defer c.Logout(5 * time.Second)
	_, err = c.StartTLS(&tls.Config{
		InsecureSkipVerify: true,
	})
	ensure.Nil(t, err)
	ensure.True(t, c.State() == imap.Login)

	// tls has worked, now login
	_, err = c.Login("test@example.com", "1234")
	ensure.Nil(t, err)
}
