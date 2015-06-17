package smtp

import (
	"crypto/tls"
	"log"

	"github.com/GeilMail/geilmail/cfg"
	"github.com/GeilMail/geilmail/helpers"
	"github.com/GeilMail/geilmail/storage/users"
)

var (
	tlsConf *tls.Config
	// localDomains is a list of domains that have mail accounts on this mail server
	localDomains []string
)

func Boot(c *cfg.Config) <-chan bool {
	var err error
	tlsConf = helpers.TLSConfig(c)

	log.Println("Booting SMTP server")
	localDomains, err = users.AllDomains()
	if err != nil {
		panic(err)
	}

	rdy := make(chan bool, 1)
	go listen(c.SMTP.ListenIP, c.SMTP.Port, c.SMTP.HostName, rdy)
	return rdy
}

func ShutDown() {
	listening = false
}
