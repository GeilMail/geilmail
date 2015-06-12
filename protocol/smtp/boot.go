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

func Boot(c *cfg.Config) <-chan bool {
	tlsConf = helpers.TLSConfig(c)
	log.Println("Booting SMTP server")
	rdy := make(chan bool, 1)
	go listen(c.SMTP.ListenIP, c.SMTP.Port, rdy)
	return rdy
}

func ShutDown() {
	listening = false
}
