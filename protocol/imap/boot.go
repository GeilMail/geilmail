package imap

import (
	"crypto/tls"
	"log"

	"github.com/GeilMail/geilmail/cfg"
	"github.com/GeilMail/geilmail/helpers"
)

var tlsConf *tls.Config

func Boot(c *cfg.Config) {
	tlsConf = helpers.TLSConfig(c)
	log.Println("Booting IMAP server")
	go listen(c.IMAP.ListenIP, c.IMAP.Port)
}
