package imap

import (
	"crypto/tls"
	"log"

	"github.com/GeilMail/geilmail/cfg"
)

var tlsConf *tls.Config

func Boot(c *cfg.Config) {
	cert, err := tls.LoadX509KeyPair(c.TLS.CertPath, c.TLS.KeyPath)
	if err != nil {
		panic(err)
	}
	tlsConf = &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	log.Println("Booting IMAP server")
	go listen(c.IMAP.ListenIP, c.IMAP.Port)
}
