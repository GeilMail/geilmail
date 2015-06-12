package smtp

import (
	"crypto/tls"
	"log"

	"github.com/GeilMail/geilmail/cfg"
	"github.com/GeilMail/geilmail/helpers"
)

var (
	tlsConf *tls.Config
)

func Boot(c *cfg.Config) {
	tlsConf = helpers.TLSConfig(c)
	log.Println("Booting SMTP server")
	go listen(c.SMTP.ListenIP, c.SMTP.Port)
}

func ShutDown() {
	listening = false
}
