package imap

import (
	"crypto/tls"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/GeilMail/geilmail/cfg"
	"github.com/facebookgo/ensure"
	"github.com/mxk/go-imap/imap"
)

const imapTestPort = 1143

func TestMain(m *testing.M) {
	gopath := os.Getenv("GOPATH")
	if string(gopath[len(gopath)-1]) == ":" {
		gopath = gopath[:len(gopath)-1]
	}

	rdy := Boot(&cfg.Config{
		IMAP: cfg.IMAPConfig{
			ListenIP: "0.0.0.0",
			Port:     imapTestPort,
		},
		TLS: cfg.TLSConfig{},
	})
	<-rdy
	ret := m.Run()
	ShutDown()
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
