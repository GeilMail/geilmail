package imap

import (
	"crypto/tls"
	"log"

	"github.com/GeilMail/geilmail/cfg"
	"github.com/GeilMail/geilmail/helpers"
)

var tlsConf *tls.Config

func Boot(c *cfg.Config) chan bool {
	tlsConf = helpers.TLSConfig(c)
	log.Println("Booting IMAP server")
	rdy := make(chan bool, 1)
	go listen(c.IMAP.ListenIP, c.IMAP.Port, rdy)
	return rdy
}
