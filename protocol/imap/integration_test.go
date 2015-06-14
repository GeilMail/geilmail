package imap

import (
	"crypto/tls"
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

func TestIMAPFlow(t *testing.T) {
	c, err := imap.Dial("localhost:1143")
	ensure.Nil(t, err)
	defer c.Logout(5 * time.Second)
	_, err = c.StartTLS(&tls.Config{
		InsecureSkipVerify: true,
	})
	ensure.Nil(t, err)
}
