package smtp

import (
	"crypto/tls"
	"log"

	"github.com/GeilMail/geilmail/cfg"
	"github.com/GeilMail/geilmail/storage/mail"
)

var (
	tlsConf     *tls.Config
	mailStorage mail.Storage
)

func Boot(c *cfg.Config) {
	cert, err := tls.LoadX509KeyPair(c.TLS.CertPath, c.TLS.KeyPath)
	if err != nil {
		panic(err)
	}

	log.Println("Booting SMTP server")
	tlsConf = &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	go listen(c.SMTP.ListenIP, c.SMTP.Port)
}

func ShutDown() {
	listening = false
}
